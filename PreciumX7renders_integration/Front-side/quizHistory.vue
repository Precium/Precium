<template>
<div class="default-container">
  <div class="default-wrap">
    <go-to-last-page-nav :title = currentPageTitle />
    <div class="points-summary">
      <div class="holding-section">
        <!-- TODO Change class naming -->
        <div class="left-holding-section">
          <div v-if="userInfo.status === 'asigned'" class="qrcode"> <!--지갑이 할당이 된 상태만 QR 코드를 보여준다.-->
            <qrcode-vue :value="walletAddress" :size="80" level="L"></qrcode-vue>
          </div>
          <div class="left-holding-section-pcm">
            <p>보유 PCM</p>
            <div class="left-holding-section-pcm-count">
              <img src="@/assets/precium_pcm.svg" alt="pcm">
              <p class="count">{{ balance | fmtPrice }}</p>
            </div>
          </div>
        </div>
        <div class="right-holding-section">
          <button v-if="userInfo.address !== undefined && userInfo.address !== 'Reserved for booking'" class="btn-line-p" @click="showWalletAddress()">지갑주소</button><br>
          <button v-if="userInfo.address !== undefined && userInfo.address !== 'Reserved for booking'" class="btn-fill-p" @click="comfirmDeposit()">입금확인</button>
        </div>
      </div>
      <span class="divider-block" />
      <div class="other-pcm">
        <div class="other-pcm-service" style="cursor: pointer;" @click="takeQuiz()">
          <p>
            퀴즈 풀기
            <img src="@/assets/Icon-right-arrow.svg" alt="이동" class="icon arrow-icon">
          </p>
        </div>
        <div class="other-pcm-service" style="cursor: pointer;" @click="withdraw()">
          <p>
            보유한 PCM <br> 출금하기
            <img src="@/assets/Icon-right-arrow.svg" alt="이동" class="icon arrow-icon">
          </p>
        </div>
        <div class="other-pcm-service" style="cursor: pointer;" @click="notice()">
          <p>
            도움말 및 <br> 유의사항
            <img src="@/assets/Icon-right-arrow.svg" alt="이동" class="icon arrow-icon">
          </p>
        </div>
      </div>
      <span class="divider-block" />
      <div class="point-usage-history">
        <div v-if="pcmUsageHistory === undefined && userInfo.status === 'asigned'" class="history-contents">
          <p class="non-usage-history">이용 내역이 없습니다</p>
        </div>
        <div v-if="pcmUsageHistory !== undefined && userInfo.status === 'asigned'" class="history-contents">
          <pcm-usage-history-list
            v-for="historyList in pcmUsageHistory" :key="historyList.id"
            :history = historyList />
        </div>
        <div v-if="userInfo.status === 'request'" class="history-contents">
          <div class="creating-wallet">
            <img src="@/assets/digital-wallet.svg" alt="wallet">
            <p class="title">새로운 지갑을 생성중입니다...</p>
            <button pclass="btn-fill-g">지갑 생성 신청 완료</button>
          </div>
        </div>
        <div v-if="userInfo.address === undefined" class="history-contents">
          <div class="create-wallet">
            <img src="@/assets/digital-wallet.svg" alt="wallet">
            <p class="title">지갑을 생성해보세요!</p>
            <button @click="createdWallet()" class="btn-fill-p">지갑 생성하기</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</div>
</template>

<script lang='js'>
import Vue from 'vue'
import Component from 'nuxt-class-component'
import { Getter, namespace } from 'vuex-class'
import GoToLastPageNav from '@/components/GoToLastPageNav'
import PcmUsageHistoryList from '@/components/PcmUsageHistoryList'
import QrcodeVue from 'qrcode.vue'
import api from '@/api'
import Swal from 'sweetalert2/dist/sweetalert2.min.js';

const AUTH = namespace('auth');

@Component({
  components: {
    GoToLastPageNav,
    PcmUsageHistoryList,
    QrcodeVue,
  },
})
export default class QuizHistory extends Vue {

  @AUTH.Getter('getIsAuthorized') getIsAuthorized;

  currentPageTitle = '내 지갑'
  userInfo = {}
  balance = ''
  pcmUsageHistory = []
  walletAddress = ''

  async created() {
    if(!this.getIsAuthorized) this.$router.push('/commerce/signin'); // 로그인 되지 않은 유저가 URL을 통해 진입 시 로그인 페이지로
    this.getUserInfo()
  }
  async getUserInfo() {
    // 유저 정보를 가저오는 API
    let res = await api.v2.market.wallet.getWalletInformations(this.$axios)
    this.userInfo = res.data
    this.userInfo.info === 'There is no wallet information' || this.userInfo.balance === undefined ? this.balance = 0 : this.balance = res.data.balance
    this.walletAddress = `ethereum:${this.userInfo.address}`

    // 사용내역에 관련된 API
    let response = await api.v2.market.wallet.walletHistory(this.$axios)
    response.data.length === 0 || response.data[0].id === 0 ? this.pcmUsageHistory = undefined : this.pcmUsageHistory = response.data
  }
  async createdWallet() {// 지갑을 생성하는 api
    let res = await api.v2.market.wallet.bindWallet(this.$axios)
    if(res.data.errorCode === 200) {
      this.$swal({
        html:
        '<p class="success-class" style="margin: 0;">지갑 생성이 완료되었습니다. 🤗</p>',
        showConfirmButton: false,
        timer: 2500,
      });
    }
    if(res.data.errorCode === 205) {
      this.$swal({
        html:
        '<p class="success-class" style="margin: 0;">지갑 신청이 완료되었습니다. 생성은 시간이 소요될 수 있습니다. 🤗</p>',
        showConfirmButton: false,
        timer: 2500,
      });
    }
    this.getUserInfo()
  }
  takeQuiz() {
    this.$router.push(`/quiz`)
  } // END take Quiz (퀴즈 페이지로 이동)
  notice() {
    this.$router.push(`/quiz/help-about-wallet`)
  } // END notice (유의사항 페이지로 이동)
  withdraw() {
    // this.$swal({
    //   html:
    //     '<p class="warning-class" style="margin: 0;">출금하기 기능은 추후 제공될 예정입니다..</p>',
    //   showConfirmButton: false,
    //   timer: 1000,
    // })
    this.$router.push('/quiz/withdraw')
  } //END withdraw (출금하기 버튼)
  showWalletAddress() {
  this.$swal({
    title: '<p class="title-class">지갑주소</p>',
    html:
        `<p class="content-text-class" style="text-align: center;">클릭해서 복사해보세요!😉</p>
        <button id="copy-url" class="copy-url-class" style="margin: 0 0 10px 0;">
        <input type="text" value="${this.userInfo.address}" id="prod-url" class="copy-url-class">
        </button>`,
      showConfirmButton: false,
      customClass: {
      },
      onBeforeOpen() {
        const content = Swal.getContent();
        const $ = content.querySelector.bind(content);
        const copyUrl = $('#copy-url');
        copyUrl.addEventListener('click', () => {
          const copyText = document.getElementById('prod-url');
          copyText.select();
          document.execCommand('copy');
          Swal.close();
          Swal.fire({
            html:
              '<p class="success-class" style="margin: 0;">복사되었습니다</p>',
            showConfirmButton: false,
            timer: 1500,
          });
        });
      },
    });
  } // END showWalletAddress (지갑주소 복사하기)
  async comfirmDeposit() {
    const loading = this.$loading({
      lock: true,
      text: '최대 수 분이 걸릴 수 있습니다. 잠시만 기다려주세요.',
      spinner: 'el-icon-loading',
      background: 'rgba(0, 0, 0, 0.7)'
    });
    let res = await api.v2.market.wallet.sweepNupdateBalance(this.$axios)
    if(res.data.errorCode === 252) {
      this.$swal({
        html:
        '<p class="warning-class" style="margin: 0;">입금한 PCM이 없습니다.</p>',
        showConfirmButton: false,
        timer: 2500,
      });
      loading.close()
      return
    }
    if(res.data.info === 'done') {
      loading.close()
      this.getUserInfo()
    }
  }
  layout(context) {
    return 'onlyBody';
  } // END layout (헤더, 풋터 제외)
}
</script>

<style lang='scss' scoped>
.points-summary {
  position: relative;
  width: 100%;
  height: 100vh;
  .holding-section {
    margin: 24px 0;
    padding: 0 16px;
    .left-holding-section {
      display: flex;
      align-items: center;
      .qrcode {
        margin-right: 18px;
        // width: 60px;
        // height: 60px;
        // background-color: #FFDDA6;
      }
      .left-holding-section-pcm {
        margin-top: 6px;
        margin-bottom: 6px;
        p:first-child {
          line-height: 24px;
        }
        .left-holding-section-pcm-count {
          display: flex;
          align-items: center;
          margin-top: 4px;
          img {
            margin-right: 6px;
            width: 22px;
            height: 22px;
          }
          .count {
            font-size: 20px;
            font-weight: bold;
          }
        }
      }
    }
    .right-holding-section {
      position: absolute;
      top: 14px;
      right: 16px;
      button {
        margin: 0;
        padding: 10px 18px;
        border-radius: 5px;
        &:first-child {
          margin: -16px 0 4px 0;
        }
        &:last-child {
          border: none;
        }
      }
    }
  }
  // .holding-points {
  //   display: flex;
  //   justify-content: space-between;
  //   align-items: center;
  //   margin: 24px 0;
  //   padding: 0 16px;
  //   .left-holding-point {
  //     .qrcode {
  //       width: 60px;
  //       height: 60px;
  //       background-color: #FFDDA6;
  //     }
  //     .aa {
  //       display: inline-block;
  //       p:first-child {
  //       line-height: 24px;
  //       }
  //       p:last-child {
  //         margin-top: 4px;
  //         font-size: 20px;
  //         font-weight: bold;
  //         line-height: 20px;
  //       }
  //     }
  //   }
  //   .redeem-point {
  //     color: #7629C4;
  //     border: 1px solid #7629C4;
  //   }
  // }
  .other-pcm {
    display: flex;
    width: 100%;
    height: 65px;
    .other-pcm-service {
      display: flex;
      flex-wrap: wrap;
      align-content: center;
      flex-basis: 33.3%;
      border-right: 1px solid #f2f2f2;
      &:last-child {
        border: none;
      }
      p {
        width: 100%;
        line-height: 16px;
        text-align: center;
        font-size: 13px;
        color: #888;
        img {
          margin: 0;
          width: 5px;
        }
      }
    }
  }
  .point-usage-history {
    width: 100%;
    // height: 100vh;
    }
    .history-contents {
      height: 100%;
      .non-usage-history {
        margin-top: 80px;
        color: #888;
        text-align: center;
      }
      .create-wallet {
        margin-top: 54px;
        padding: 0 16px;
        img {
          display: block;
          margin: 0 auto;
          width: 74px;
          height: 74px;
        }
        .title {
          padding-top: 16px;
          padding-bottom: 24px;
          font-size: 20px;
          font-weight: bold;
          text-align: center;
        }
        button {
          padding: 12px 0;
          width: 100%;
          border: none;
          border-radius: 5px;
          font-size: 16px;
        }
      }
      .creating-wallet {
        margin-top: 54px;
        padding: 0 16px;
        img {
          display: block;
          margin: 0 auto;
          width: 74px;
          height: 74px;
        }
        .title {
          padding-top: 16px;
          padding-bottom: 24px;
          font-size: 20px;
          font-weight: bold;
          text-align: center;
        }
        button {
          padding: 12px 0;
          width: 100%;
          color: #888;
          border: none;
          border-radius: 5px;
          font-size: 16px;
        }
      }
    }
  .arrow-icon {
    margin-left: 4px;
    // NOTE change svg color black to grayscale
    filter: brightness(0) invert(0.5);
  }
}
</style>

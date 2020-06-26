<template>
<div class='withdraw-index default-container'>
  <div class="default-wrap">
    <go-to-last-page-nav :title = currentPageTitle />
    <div class="withdraw-title">
      <img src="@/assets/digital-wallet.svg" alt="wallet">
      <p>보유한 PCM 출금하기</p>
    </div>
    <span class="divider-block" />
    <div class="withdraw-services">
      <div @click="showWalletModal()" class="add-withdraw-wallet">
        <p class="services-btn">출금 지갑 등록</p>
      </div>
      <no-ssr>
        <modal
          class="add-wallet-modal"
          name="showWalletModal"
          width="350"
          height="460"
        >
          <div class="modal-title">
            <span>출금 지갑 등록</span>
            <img src="~/assets/close.svg" alt="닫기" @click="$modal.hide('showWalletModal')" >
          </div>
          <div class="modal-input">
            <p>지갑 별칭 (alias)</p>
            <!-- NOTE 최소 1자 최대 10자 제한 걸어줘야함 -->
            <input
              type="text"
              class="wallet-alias"
              v-model="walletList.alias"
              v-validate="'required|min:1|max:10'"
              placeholder="지갑 별칭을 정해주세요 (최소 1자 최대 10자)">
            <p>출금 지갑 주소</p>
            <input v-model="walletList.address" type="text" class="wallet-address" placeholder="ex) 0x54e...">
          </div>
          <div class="modal-checkbox">
            <div class="check">
              <input type="checkbox" name="check1" id="check1">
              <label for="check1">
                <span/><p>잘못 기입한 출금 지갑 주소에 관한 책임 소지는 전적으로 본인에게 있음에 동의합니다.</p>
              </label>
            </div>
            <div class="check">
              <input type="checkbox" name="check2" id="check2">
              <label for="check2">
                <span/><p>'출금하기 유의사항'에 관해 정독하였으며, 해당 내용에 전부 동의합니다.</p>
              </label>
            </div>
          </div>
          <button @click="addNewWithdrawWallet()" class="btns btn-long btn-long-fill-p">등록하기</button>
        </modal>
      </no-ssr>
      <div @click="isWithdrawPW()" class="set-withdraw-password">
        <p class="services-btn">출금 비밀번호</p>
      </div>
      <no-ssr>
        <modal
          class="set-password-modal"
          name="showPasswordModal"
          width="350"
          height="485"
        >
          <div class="modal-title">
            <span>출금 비밀번호 설정</span>
            <img src="~/assets/close.svg" alt="닫기" @click="$modal.hide('showPasswordModal')" >
          </div>
          <div class="modal-input">
            <p>출금 비밀번호</p>
            <!-- NOTE 숫자 6자리 제한 걸어줘야함 -->
            <input
              type="password"
              class="enter-password"
              v-model="enterPW"
              v-validate="'required|numeric|min:1|max:6'"
              pattern="[0-9]*"
              inputmode="numeric"
              placeholder="비밀번호를 입력해주세요 (숫자 6자리)"
            >
            <p>비밀번호 재입력</p>
            <input
              type="password"
              class="re-password"
              v-model="reEnterPW"
              v-validate="'required|numeric|max:6'"
              pattern="[0-9]*"
              inputmode="numeric"
              placeholder="비밀번호를 재입력해주세요"
            >
            <div class="phone-confirm">
              <p>번호 인증</p>
              <div>
                <input
                  type="number"
                  class="phone-number"
                  v-validate="'required|numeric|min:10|max:11'"
                  pattern="[0-9]*"
                  placeholder="휴대폰번호 (- 없이 입력하세요)"
                >
                <button class="btn-line-g">인증번호발송</button>
              </div>
              <div>
                <input
                  type="number"
                  class="confirm-number"
                  pattern="[0-9]*"
                  inputmode="numeric"
                  placeholder="인증번호"
                >
                <span>03:00</span>
                <button class="btn-line-p">인증하기</button>
              </div>
            </div>
          </div>
          <button
            @click="setNewPassword()"
            class="btns btn-long btn-long-fill-p">등록하기</button>
        </modal>
      </no-ssr>
      <div class="withdraw-request-history" @click="goToRequestHistory">
        <p class="services-btn">출금 요청 내역</p>
      </div>
    </div>
    <span class="divider-block" />
    <div class="withdraw-wallet-list">
      <div v-if="this.walletList.length === 0" class="no-wallet empty-list">
        <p>🤔</p>
        <p>등록된 출금 지갑이 없어요</p>
      </div>
      <div v-else class="have-wallet">
        <withdraw-wallet-list v-for="withdrawWallet in walletList"
          :key="withdrawWallet.id"
          :wallet = withdrawWallet
          @showWithdrawWallet="requestWithdrawModal"
          @deleteWalletList="onClickDeleteWallet"
        />
        <no-ssr>
          <modal
            class="request-withdraw"
            name="requestWithdrawModal"
            width="350"
            height="440"
          >
            <div class="modal-title">
              <span>출금하기</span>
              <img src="~/assets/close.svg" alt="닫기" @click="$modal.hide('requestWithdrawModal')" >
            </div>
            <div class="modal-input">
              <p>출금 요청 지갑</p>
              <input style="background-color: #f2f2f2; color: #888;" readonly="readonly">
              <p>출금 요청 PCM</p>
              <input
                type="number"
                v-validate="'required|numeric|max:6'"
                placeholder="0"
              >
              <p>출금 비밀번호</p>
              <input
                type="password"
                placeholder="출금 비밀번호를 입력해주세요"
              >
            </div>
            <button class="btns btn-long btn-long-fill-p">출금 요청하기</button>
          </modal>
        </no-ssr>
      </div>
    </div>
    <span class="divider-block" />
    <div class="notice cursor-pointer" @click="active = !active">
      <span class="notice-title">출금하기 유의사항</span>
      <img src="@/assets/down-arrow.svg" :class="{'arrow180':active}"/>
    </div>
    <vue-slide-up-down class="notice-content" :active="active">
    </vue-slide-up-down>
    <span class="divider-block" />
  </div>
</div>
</template>

<script lang='js'>
import Vue from 'vue'
import Component from 'nuxt-class-component'
// import api from '@/api'
// import { Getter, Action, namespace } from 'vuex-class';
// import SHA256 from 'crypto-js/sha256';
import GoToLastPageNav from '@/components/GoToLastPageNav'
import WithdrawWalletList from '@/components/WithdrawWalletList'

@Component({
  components: {
    GoToLastPageNav,
    WithdrawWalletList,
  },
})
export default class withdraw extends Vue {
  // vuex
  // @AUTH.Getter('getIsAuthorized') getIsAuthorized;

  currentPageTitle = '출금하기'
  active = false //vue-slide-up-down
  
  //임시데이터
  enterPW = ''
  reEnterPW = ''
  isPassword = false // 출금비밀번호 등록여부
  walletList = []

  created() {
    /* 
      유저 정보 필요함
      - 아이디
      - 출금지갑리스트 this.refreshWalletList()
      - 보유한 PCM this.holdingPCM()
    */
    console.log(this.walletList)
  }

  showWalletModal() { // 출금지갑등록 modal
    if(this.walletList.length >= 3) {
      this.$swal({
        html:
          '<p class="warning-class">🚫 출금 지갑은 최대 3개까지만 등록 가능합니다!</p>',
        showConfirmButton: false,
        timer: 1500
      })
      return
    } else if(0 <= this.walletList.length < 3) {
      this.walletList.alias = ''
      this.walletList.address = ''
      this.$modal.show('showWalletModal')
    } else {
      this.$swal({
        html: '<p class="warning-class">현재 출금 지갑을 등록할 수 없습니다.</p>',
        showConfirmButton: false,
        timer: 1500
      })
    }
  }

  addNewWithdrawWallet() { // 출금지갑 등록하기 btn
    if(this.walletList.alias === '' || this.walletList.address === '') {
      this.$swal({
        html:
          '<p class="warning-class">빈 항목을 모두 채워주세요!</p>',
        showConfirmButton: false,
        timer: 1500,
      })
      return
    } if(this.walletList.alias.length < 1 || this.walletList.alias.length > 10) {
      this.$swal({
        html:
          '<p class="warning-class">지갑 별칭은 최소1자 ~ 최대10자로 설정해주세요!</p>',
        showConfirmButton: false,
        timer: 1500,
      })
      return
    }
    else {
      this.$swal({
        html:
          '<p class="success-class">새로운 출금 지갑을 등록했습니다!</p>',
        showConfirmButton: false,
        timer: 1500,
      })
      this.$modal.hide('showWalletModal')
      // NOTE api 통신을 통해 사라질 push부분
      // FIXME this.refreshWalletList()
      this.walletList.push({
        alias: this.walletList.alias,
        address: this.walletList.address
      })
      // this.refreshWalletList()
      this.isWithdrawPW()
    }
    console.log(this.walletList)
  }

  refreshWalletList() {
    // 1. api통신 -> 현재 등록된 유저의 출금 지갑 리스트 가져옴
    // 2. 가져온 지갑 리스트 데이터를 this.wallet.List에 바인딩
    // let walletInfo = api.v2.wallet...
  }

  isWithdrawPW() {
    // 1. api통신 -> 유저의 출금 비밀번호 등록 유무를 판단
    // 2. 출금 비밀번호가 등록되지 않은 경우 비밀번호 등록 진행 showPasswordModal()
    // 3. 출금 비밀번호가 등록된 경우 재설정 swal 띄운 후 than으로 showPasswordModal()
    if(this.isPassword) {
      this.$swal({
        html:
          '<p class="success-class">비밀번호를 재설정 하시겠습니까?</p>',
        showCancelButton: true,
        cancelButtonText: '취소',
        confirmButtonText: '재설정',
        confirmButtonClass: 'confirm-btn-class',
        cancelButtonClass: 'cancel-btn-class',
        reverseButtons: true,
      }) .then((res => {
        if (res.value) {
          this.$modal.show('showPasswordModal')
        }
      }))
    } else {
      this.$modal.show('showPasswordModal')
    }
  }

  holdingPCM() {
    // 1. api통신 -> 유저가 보유한 PCM 개수를 가져옮
    // 2. 출금 요청 시 출금 요청 PCM validate 체크할 때 사용
  }

  setNewPassword() {
    // 1. api통신 -> 출금 비밀번호 최초설정/재설정 시 서버에 비밀번호 업데이트
    // 2. SHA256으로 비번을 감싸서 해쉬값을 넘겨야 함
    if (this.enterPW === '' || this.reEnterPW === '') {
      this.$swal({
        html:
          '<p class="warning-class">빈 항목을 모두 채워주세요!</p>',
        showConfirmButton: false,
        timer: 1500,
      })
      return
    }
    if (this.enterPW !== this.reEnterPW) {
      this.$swal({
        html:
          '<p class="warning-class">비밀번호가 일치하지 않습니다!</p>',
        showConfirmButton: false,
        timer: 1500,
      })
      return
    }
    // if ()
    this.$swal({
      html:
        '<p class="success-class">새 비밀번호를 설정했습니다.</p>',
      showConfirmButton: false,
      timer: 1500
    })
    console.log(this.isPassword)
    this.isPassword = true
    console.log(this.isPassword)
    this.$modal.hide('showPasswordModal')
  }

  requestWithdrawModal() {
    if(this.isPassword) {
      this.$modal.show('requestWithdrawModal')
    } else {
      this.$swal({
        html:
          '<p class="warning-class">출금 비밀번호를 먼저 설정해주세요!</p>',
        showConfirmButton: false,
        timer: 1500
      })
      this.$modal.show('showPasswordModal')
    }
  }

  onClickDeleteWallet(walletID) {
    // 1. api통신 -> walletList 중 walletID에 해당하는 지갑을 뺀다고 업데이트
    // 2. NOTE 뺀다고 업데이트 이후 바로 refreshWalletList() ?
    // 3. NOTE or 어차피 서버에 업데이트 됐기 때문에 front단에서 컴포넌트 삭제된 것처럼만 처리하고 나중에 이 페이지가 새로고침 되면 created()에서 refreshWalletList()?
    console.log('컴포넌트에서 받아온 wallet id')
    console.log(walletID)
    let deleteWallet = this.walletList.splice(walletID)
    console.log('ㅡㅡㅡㅡㅡㅡㅡ')
    console.log(deleteWallet)
    // return
    this.walletList = deleteWallet
    console.log('새로운지갑배열')
    console.log(this.walletList)
  }

  goToRequestHistory() {
    this.$router.push('/quiz/withdraw-request')
  }

  layout(context) {
    return 'onlyBody';
  }
}
</script>

<style lang='scss' scoped>
.withdraw-index {
  width: 100%;
  height: 100vh;
  .withdraw-title {
    padding: 24px;
    display: flex;
    align-items: center;
    p {
      padding-left: 12px;
      font-size: 20px;
      font-weight: bold;
    }
  }
  .withdraw-services {
    display: flex;
    width: 100%;
    height: 64px;
    div {
      display: flex;
      align-content: center;
      flex-basis: 33.3%;
      border-right: 1px solid #f2f2f2;
      &:last-child {
        border: none;
      }
    }
    p.services-btn {
      width: 100%;
      color: #888;
      font-size: 14px;
      line-height: 64px;
      text-align: center;
    }
  }
  .withdraw-wallet-list {
    width: 100%;
    .no-wallet {
      width: 100%;
      max-width: 640px;
      height: 300px;
      margin: 0 auto;
      padding: 32px 16px 0;
      text-align: center;
      p:first-child {
        margin-top: 60px;
        font-size: 50px;
        opacity: 0.4;
      }
      p:last-child {
        margin-top: 24px;
        color: #888;
        font-size: 24px;
        font-weight: bold;
      }
    }
    .have-wallet {
      width: 100%;
      margin-top: 24px;
      padding: 0 16px;
      .withdraw-wallet-list {
        margin: 16px 0;
      }
    }
  }
  .notice {
    position: relative;
    // margin-bottom: -4px;
    padding: 16px;
    border: none;
    span {
      width: 90%;
      font-weight: bold;
      line-height: 20px;
    }
    img {
      position: absolute;
      top: 50%;
      right: 20px;
      width: 15px;
      margin-top: -7.5px;
    }
  }
  .notice-content {
    height: auto;
    background-color: #fafafa;
    p {
      padding: 24px 16px;
      font-size: 13px;
      line-height: 20px;
      color: #888;
      word-break: keep-all;
      white-space: pre-line;
    }
  }
  /* modal css START */
  .add-wallet-modal,
  .set-password-modal,
  .request-withdraw {
    .modal-title {
      span {
        position: relative;
        display: block;
        width: 100%;
        height: 54px;
        text-align: center;
        line-height: 54px;
        border-bottom: 1px solid #f2f2f2;
      }
      img {
        position: absolute;
        right: 16px;
        top: 18px;
        cursor: pointer;
      }
    }
    .modal-input {
      display: flex;
      flex-wrap: wrap;
      p {
        width: 100%;
        margin: 16px 0 8px 16px;
        font-size: 14px;
        line-height: 20px;
        text-align: left;
        color: #222;
      }
      input {
        display: block;
        width: 320px;
        height: 50px;
        margin: 0 auto;
        padding-left: 16px;
        border: 1px solid #f2f2f2;
        border-radius: 5px;
      }
    }
    .modal-checkbox {
      display: flex;
      flex-wrap: wrap;
      width: 100%;
      margin-top: 20px;
      padding: 0 16px;
      .check {
        display: flex;
        flex-basis: 100%;
        width: 100% !important;
        border: 0;
        label {
          align-items: flex-start;
          margin-bottom: 8px;
          font-size: 14px;
          line-height: 22px;
          span {
            margin-top: 4px;
            margin-right: 12px;
          }
          p {
            width: 90%;
            margin: 0;
          }
        }
      }
    }
    .phone-confirm {
      display: flex;
      flex-wrap: wrap;
      div {
        position: relative;
        display: flex;
        justify-content: space-between;
        width: 330px;
        padding-left: 16px;
        &:last-child {
          margin-top: 8px;
        }
        span {
          position: absolute;
          top: 16px;
          right: 110px;
          color: #d93f3f;
        }
        button {
          flex-basis: 45%;
          border-radius: 5px;
          cursor: pointer;
        }
      }
    }
    button.btns {
      display: block;
      width: 90%;
      margin: 24px auto 0;
    }
  }
  /* modal css END */
}
.arrow180 {
  transform: rotate(180deg);
}
</style>

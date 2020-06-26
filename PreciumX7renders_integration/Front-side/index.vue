<template>
<div id="go-to-top" class='quiz-index default-container'>
  <div class="default-wrap">
    <go-to-last-page-nav :title = currentPageTitle />
    <div class="quiz-image">
      <img :src='quizDetail.imageURI' alt="퀴즈">
      <div v-if="quizDetail.status === 'done'" class="end-quiz-image">
        <p>종료된 퀴즈입니다.</p>
      </div>
    </div>
    <div v-if="quizDetail.status === 'open'" class="correct-answer">
      <div class="correct-answer-input">
        <input v-model="answerInput" @keyup.enter="submitAnswer()" type="text" placeholder="정답을 입력해주세요.">
        <button @click="submitAnswer()">제출하기</button>
      </div>
      <div class="holding-pcm">
        <p class="title">보유한 PCM</p>
        <div class="icon" />
        <p class="count">{{ holdingPCM | fmtFloor }}</p>
      </div>
    </div>
    <div v-if="quizDetail.status === 'done'">
      <div class="end-result">
        <div class="end-quiz"
          v-if="quizDetail.isParticipant === true && quizDetail.isWon === undefined">
          <p class="end-quiz-result btn-fill-b">다음 퀴즈에 도전해보세요😢</p>
          <p class="end-quiz-info">* 아쉽게도 회원님은 이번 퀴즈에 당첨되지 못했습니다.</p>
        </div>
        <div class="end-quiz"
          v-if="quizDetail.isParticipant === true && quizDetail.isWon === true">
          <p class="end-quiz-result btn-fill-p cursor-pointer" @click="receiveReward()">
            축하합니다! 당첨금을 환급 받으세요!🎉
          </p>
          <p class="end-quiz-info">* 오직 위 환급 버튼을 통해서만 환급이 가능합니다.</p>
        </div>
      </div>
    </div>
    <div class="quiz-information">
      <!-- 아래 '이번 퀴즈에 쌓인 PCM' 차후 되살릴 수 있음 -->
      <!-- <div class="quiz-information-condition">
        <div class="quiz-condition">
          <div class="icon">
            <img src="@/assets/precium_pcm.svg" alt="pcm">
          </div>
          <p class="title">이번 퀴즈에 쌓인 PCM</p>
          <p class="sub-title">{{ quizDetail.accumulatedPCM | fmtPrice }}PCM</p>
        </div>
      </div> -->
      <div class="quiz-information-condition">
        <div class="quiz-condition">
          <div class="icon">
            <img src="@/assets/precium_pcm.svg" alt="pcm">
          </div>
          <p class="title">이번 퀴즈 참여 조건</p>
          <p class="sub-title">{{ quizDetail.requiredPCM | fmtFloor }}PCM</p>
        </div>
      </div>
      <div class="quiz-information-condition">
        <div class="quiz-condition">
          <div class="icon">
            <img
              src="@/assets/profile.svg"
              alt="number"
              style="width:54%;height:54%;margin-top:6px;">
          </div>
          <p class="title">이번 퀴즈 유저 총 참여 횟수</p>
          <p class="sub-title">{{ quizDetail.participantsNumber | fmtPrice }}회</p>
        </div>
      </div>
      <div class="quiz-information-condition">
        <div class="quiz-condition">
          <div class="icon">
            <img src="@/assets/precium_pcm.svg" alt="pcm">
          </div>
          <p class="title">분배율</p>
          <p class="sub-title">기부 40% 당첨자 40% 플랫폼 20%</p>
        </div>
      </div>
      <div class="quiz-information-condition">
        <div class="quiz-condition">
          <div class="icon">
            <img
              src="@/assets/profile.svg"
              alt="number"
              style="width:54%;height:54%;margin-top:6px;">
          </div>
          <p class="title">퀴즈 당첨자</p>
          <p class="sub-title">정답자 중 총 20명 예정</p>
        </div>
      </div>
    </div>
    <div class="participation">
      <p class="participation-title">참여방법</p>
      <div class="participation-means">
        <div class="participation-means-step">
          <div class="icon">
            <img src="@/assets/precium_pcm.svg" alt="pcm">
          </div>
          <p class="title">STEP 1</p>
          <p class="sub-title">PCM 충전</p>
        </div>
        <div class="participation-means-step">
          <div class="icon">
            <img src="@/assets/submit-quiz.svg" alt="submit">
          </div>
          <p class="title">STEP 2</p>
          <p class="sub-title">정답 제출</p>
        </div>
        <div class="participation-means-step">
          <div class="icon">
            <img src="@/assets/iswon-quiz.svg" alt="won">
          </div>
          <p class="title">STEP 3</p>
          <p class="sub-title">퀴즈 종료 후 당첨 확인</p>
        </div>
      </div>
    </div>
    <span class="divider-block" />
    <div class="last-quiz">
      <p class="last-quiz-title">지난퀴즈 🕹</p>
        <prev-quiz
        v-for="prevList in prevQuizList" :key="prevList.id"
        :quizList = prevList
        @sendPrevQuizDetail="modifyQuizDetail" 
        v-scroll-to="'#go-to-top'"/>
    </div>
    <span class="divider-block" />
    <div class="notice-list">
      <div class="notice" @click="active = !active">
        <span class="notice-title">퀴즈 참여 시 유의사항</span>
        <img src="@/assets/down-arrow.svg" alt="" :class="{'arrow180':active}"/>
      </div>
      <vue-slide-up-down class="notice-answer" :active="active">
        <div class="notice-content">
          <p>- 트렌더스 퀴즈는 트렌더스 회원이라면 누구나 퀴즈에 참여할 수 있는 서비스입니다. 단, 지갑이 생성되어 있어야 참여가 가능합니다.</p>
          <p>- 트렌더스 퀴즈에 게시된 문제의 정답을 맞히기 위해선 보유한 PCM이 있어야 하며, 퀴즈가 진행되는 동안 참여할 수 있는 PCM 개수는 고정입니다.</p>
          <p>- 사용자가 참여해서 쌓인 PCM 중 40%를 정답을 맞힌 정답자 중 20명에게 랜덤으로 지급될 예정입니다.</p>
          <p>- 퀴즈를 푸는 도중 정답을 못 맞히더라도 재참여가 가능하며, 정답을 맞힌 참여자도 재참여가 가능합니다. 따라서, 여러 번 정답을 맞힐수록 당첨 확률이 올라갑니다.</p>
          <p>- 퀴즈 생성은 주 1회 한 번씩 생성(주말 제외) 됩니다. 3일 동안 퀴즈에 참여 가능하고 2일은 트렌더스 측에서 정산 후 PCM 지급이 진행됩니다.</p>
          <p>- 퀴즈 결과는 퀴즈가 종료되고 해당 퀴즈 페이지에서 확인이 가능하며, 퀴즈에 당첨된 사람에게 PCM이 지급됩니다. 당첨자는 '당첨금 환급' 버튼을 클릭해야 PCM이 지급되고 환급받은 PCM 확인은 '내 지갑' 메뉴를 통해 확인이 가능합니다.</p>
          <button class="btn-fill-b" @click="goToHelpWallet()">
          지갑 생성 도움말 보기
          </button> 
        </div>
      </vue-slide-up-down>
    </div>
    <span class="divider-block" />
  </div>
</div>
</template>

<script lang='js'>
import Vue from 'vue'
import Component from 'nuxt-class-component'
import api from '@/api'
import { Getter, Action, namespace } from 'vuex-class';
import GoToLastPageNav from '@/components/GoToLastPageNav'
import PrevQuiz from '@/components/PrevQuiz'
import VueScrollTo from 'vue-scrollto'
import { fmtPrice, fmtDate, fmtFloor } from '@/plugins/filter'

const AUTH = namespace('auth');

@Component({
  components: {
    GoToLastPageNav,
    PrevQuiz,
  },
})
export default class quizList extends Vue {
  // vuex
  @AUTH.Getter('getIsAuthorized') getIsAuthorized;

  currentPageTitle = '퀴즈'
  viralUID = ''
  holdingPCM = '' // 보유한 PCM
  answerInput = '' // 제출한 정답
  correct = '' // 퀴즈 정답
  active = false // 유의사항 열고 닫기
  prevQuizList = '' // 지난 퀴즈 리스트 (배열)
  isCorrect = '' // 정답 여부
  quizDetail = {
    id: '',
    title: '', 
    content: '', 
    imageURI: '', 
    accumulatedPCM: '', 
    participantsNumber: '', 
    requiredPCM: '', 
    startAt: '', 
    endAt: '', 
    status: '', // done, open
    correct: '',
    reward: '', 
    isParticipant: '', // false
    isWon: '', // false
  }

  async created() {
    // 유저 정보를 가져오는 API
    this.getQuizInfo()
    this.refreshHoldingPCM()
  }

  // 초기값 및 컴포넌트 클릭 시 변경되는 값 공통으로 뺌
  refreshView(obj) {
    this.quizDetail.id = obj.id
    this.quizDetail.title = obj.title
    this.quizDetail.content = obj.content
    this.quizDetail.imageURI = obj.image_uri
    this.quizDetail.accumulatedPCM = obj.participants_amount
    this.quizDetail.participantsNumber = obj.participants_number
    this.quizDetail.requiredPCM = obj.consume_amount
    this.quizDetail.startAt = obj.start_at
    this.quizDetail.endAt = obj.end_at
    this.quizDetail.status = obj.status
    this.quizDetail.correct = obj.correct
    this.quizDetail.reward = obj.reward_per_user
    this.quizDetail.isParticipant = obj.is_participant
    this.quizDetail.isWon = obj.is_won
  }

  // 유저의 정답 제출을 서버로 보내 정답 유무 판단
  async sendParams() {
    let params = {
      qid: this.quizDetail.id,
      answer: this.answerInput,
      request_amount: this.quizDetail.requiredPCM
    };
    try{
      let res = await api.v2.market.wallet.quizConfirmation(this.$axios, params)
      this.isCorrect = res.data.result
      if(this.isCorrect) { // 정답인 경우
        this.$swal({
          html:
            `<p class="success-class" style="margin-bottom: 0 !important;">정답입니다!🥳</p><br>
            <p class="content-text-class" style="text-align:center !important;">
            *퀴즈 종료 후 정답자에 한 해 당첨이 결정됩니다.<br>
            *정답을 여러번 맞출 수록 당첨 확률이 올라갑니다.<br>
            *자세한 사항은 <span class="warning-class">'유의사항'</span>을 참고해주세요.</p>`,
          showConfirmButton: false,
          timer: 4000,
        })
      }else{ // 오답인 경우
        this.$swal({
          html:
            `<p class="warning-class" style="margin-bottom: 0 !important;">오답입니다😱 다시 도전해보세요!</p><br>
            <p class="content-text-class" style="text-align:center !important;">
            *퀴즈 종료 후 정답자에 한 해 당첨이 결정됩니다.<br>
            *정답을 여러번 맞출 수록 당첨 확률이 올라갑니다.<br>
            *자세한 사항은 <span class="warning-class">'유의사항'</span>을 참고해주세요.</p>`,
          showConfirmButton: false,
          timer: 4000,
        })
      }
      this.refreshHoldingPCM()
      this.getQuizInfo()
      this.answerInput = ''
    }catch(error) {
      if(error.response.data.errorCode === '130') {
        this.$swal({
          html:
            `<p class="warning-class">
            PCM 시세 변동으로 인해 사용이 어렵습니다.<br>
            보유 PCM과 시세를 체크해주세요!</p>`,
          showConfirmButton: false,
          timer: 2500,
        })
      }else if (error.response.data.errorCode === '120'){
        this.$swal({
          html:
            `<p class="warning-class">
            PCM이 부족합니다! PCM을 충전해주세요.</p>`,
          showConfirmButton: false,
          timer: 2000,
        })
      }
    }
  }

  submitAnswer() {
    if(this.getIsAuthorized) { // 회원일 경우
      if(this.answerInput === '') { // 정답란이 빈칸인 경우
        this.$swal({
          html:
            `<p class="warning-class">정답을 입력해주세요!</p>`,
          showConfirmButton: false,
          timer: 1000,
        })
        return
      }
      // 정답란이 빈칸이 아닌 경우 + 참여 조건이 만족할 경우
      if(this.quizDetail.requiredPCM <= this.holdingPCM) {
        this.$swal({
          html:
            `<p class="title-class">PCM을 사용합니다. 정답을 제출할까요?</p>`,
          showCancelButton: true,
          confirmButtonText: '제출하기',
          cancelButtonText: '취소',
          confirmButtonClass: 'confirm-btn-class',
          cancelButtonClass: 'cancel-btn-class',
          reverseButtons: true,
        }).then((res) => {
          if (res.value) {
            if(this.quizDetail.requiredPCM <= this.holdingPCM) {
              this.sendParams()
            }
          }
        })
      }else { // 정답란이 빈칸이 아닌 경우 + 참여 조건이 부족할 경우
        this.$swal({
          html:
            `<p class="title-class">PCM이 부족합니다. 충전하시겠습니까?</p>`,
          showCancelButton: true,
          confirmButtonText: '충전하기',
          cancelButtonText: '취소',
          confirmButtonClass: 'confirm-btn-class',
          cancelButtonClass: 'cancel-btn-class',
          reverseButtons: true,
        }).then((res) => {
          if (res.value) {
            this.$router.push('/quiz/quizHistory')
          }
        })
      }
    }else { // 비회원일 경우
      this.$router.push('/commerce/signin')
    }
  }

  async receiveReward() { 
    let rewardInfo = await api.v2.market.wallet.quizObtainReward(this.$axios, this.quizDetail.id)
    if(rewardInfo.data.info === 'done') { // 당첨금을 환급 받는 경우
      this.$swal({
        html:
          `<p style="font-weight: 500;">당첨금 ${this.quizDetail.reward}PCM을 환급 받았습니다!👏</p>`,
        showConfirmButton: false,
        timer: 3000
      })
      this.refreshHoldingPCM()
    }else {
      this.$swal({
        html:
          `<p class="warning-class">
          당첨 여부를 확인할 수 없습니다🤔</p>`,
        showConfirmButton: false,
        timer: 1500
      })
    }
  }

  async refreshHoldingPCM() {
    let balanceInfo = undefined
    // 회원일 경우 wallet balance를 가져옴
    if (this.getIsAuthorized) {
      balanceInfo = await api.v2.market.wallet.getWalletInformations(this.$axios)
      console.log(`balanceInfo =`)
      console.log(balanceInfo)
    }
    // balance를 가져와서 holdingPCM 으로 치환
    if (balanceInfo !== undefined ){
      this.holdingPCM = balanceInfo.data.balance
    }
    // 정보를 가져왔지만 balance가 존재하지 않는 경우(지갑에 PCM이 존재하지 않는 경우)
    if (balanceInfo.data.balance === undefined){
      this.holdingPCM = 0
    }
  }

  async modifyQuizDetail(target) { 
    // PrevQuiz 컴포넌트 클릭 시 $emit을 통해 실행
    let res = await api.v2.market.wallet.quizDetail(this.$axios, target.id, this.getIsAuthorized)
    // 받아온 퀴즈 리스트 값으로 refreshView
    target.participants_amount = res.data.total
    target.is_participant = res.data.is_participant
    target.is_won = res.data.is_won
    target.consume_amount = res.data.consume_amount
    target.reward_per_user = res.data.reward_per_user
    this.refreshView(target)
    // refreshView 이후 최상단으로 이동
    VueScrollTo.scrollTo(`#selectID`, { offset: -100});
  }
  async getQuizInfo() {
    let listInfo = await api.v2.market.wallet.quizList(this.$axios)
    let detailInfo = await api.v2.market.wallet.quizDetail(this.$axios, listInfo.data[0].id, this.getIsAuthorized)

    // PrevQuiz 컴포넌트를 위함
    this.refreshView(listInfo.data[0])
    this.prevQuizList = listInfo.data
    this.quizDetail.accumulatedPCM = detailInfo.data.total
    this.quizDetail.requiredPCM = detailInfo.data.consume_amount
    this.quizDetail.isParticipant = detailInfo.data.is_participant 
    this.quizDetail.isWon = detailInfo.data.is_won
    this.quizDetail.reward = detailInfo.data.reward_per_user
  }

  goToHelpWallet() {
    this.$router.push('/quiz/help-about-wallet')
  }

  layout(context) {
    return 'onlyBody';
  }
}
</script>

<style lang='scss' scoped>
.default-wrap {
  position: relative;
}
.quiz-image {
  position: relative;
  width: 100%;
  text-align: center;
  img {
    width: 100%;
    height: 100%;
  }
  .end-quiz-image {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    background: rgba(34, 34, 34, 0.5);
    p {
      padding: 8px 20px;
      border: 5px solid #fff;
      font-size: 30px;
      font-weight: 600;
      color: #fff;
      text-align: center;
    }
  } 
}
.correct-answer {
  text-align: center;
  background-color: #f2f2f2;
}
.correct-answer-input {
  display: inline-block;
  margin-right: 5%;
  margin-left: 5%;
  width: 90%;
  padding-top: 40px;
  padding-bottom: 16px;
  input {
    float: left;
    width: 70%;
    height: 50px;
    padding-left: 16px;
    line-height: 50px;
    border: 1px solid #ddd;
    border-radius: 5px;
    font-size: 16px;
  }
  input::placeholder {
    font-size: 16px;
    color: #ddd;
  }
  button {
    float: right;
    width: 28%;
    height: 48px;
    line-height: 48px;
    border: none;
    border-radius: 5px;
    background-color: #222;
    font-size: 16px;
    color: #fff;
  }
}
.end-result {
  background-color: #f2f2f2;
}
.end-quiz {
  margin-right: 5%;
  margin-left: 5%;
  width: 90%;
  padding-top: 40px;
  padding-bottom: 40px;
  .end-quiz-result {
    height: 50px;
    line-height: 50px;
    border-radius: 5px;
  }
  .end-quiz-info {
    padding-top: 16px;
    font-size: 13px;
    text-align: center;
    color: #222;
  }
}
.holding-pcm {
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 0 30px 40px 30px;
  .title {
    margin-right: 10px;
    font-size: 14px;
    font-weight: 600;
    color: #222;
  }
  .icon {
    margin-right: 4px;
    width: 20px;
    height: 20px;
    background: url('/_nuxt/assets/precium_pcm.svg');
    background-size: 100%;
  }
}
.quiz-information {
  border-bottom: 10px solid #f2f2f2;
  .quiz-information-condition {
    border-bottom: 1px solid #f2f2f2;
    .quiz-condition {
      display: flex;
      align-items: center;
      margin: 0 auto;
      width: 300px;
      height: 56px;
      line-height: 56px;
      // background-color: red;
      .icon {
        margin: 0 10px;
        // margin-right: 10px;
        width: 32px;
        height: 32px;
        border-radius: 33px;
        background-color: #f2f2f2;
        img {
          display: block;
          width: 65%;
          height: 65%;
          margin: 6px auto 0;
        }
      }
      .title {
        padding-right: 10px;
        font-size: 12px;
        font-weight: 600;
        color: #222;
      }
      .sub-title {
        font-size: 12px;
        color: #222;
      }
    }
  }
}
.participation {
  padding: 32px 0;
  margin-right: 5%;
  margin-left: 5%;
  width: 90%;
  .participation-title {
    padding-bottom: 20px;
    font-size: 20px;
    font-weight: 600;
    color: #222;
  }
  .participation-means {
    display: flex;
    align-self: center;
    text-align: center;
    .participation-means-step {
      flex: 3;
      .icon {
      margin: 0 auto;
      width: 80px;
      height: 80px;
      border: 1px solid #ddd;
      border-radius: 81px;
        img {
          display: block;
          width: 60%;
          height: 60%;
          margin: 16px auto;
        }
      }
      .title {
        padding-top: 8px;
        font-size: 14px;
        font-weight: 600;
        line-height: 20px;
        color: #222;
      }
      .sub-title {
        width: 80px;
        margin: 0 auto;
        font-size: 14px;
        line-height: 18px;
        color: #222;
      }
    }
  }   
}
.notice-list {
  .notice {
    position: relative;
    height: auto;
    margin-top: -1px;
    padding: 16px;
    border: none;
    cursor: pointer;
    span {
      display: block;
      width: 90%;
      margin-left: 1%;
      font-size: 16px;
      font-weight: 600;
    }
    img {
      position: absolute;
      margin-top: -7.5px;
      top: 50%;
      right: 20px;
      width: 15px;
    }
  }
  .notice-answer {
    height: auto;
    margin-top: -1px;
    border-top: none;
    background-color: #FAFAFA;
    .notice-content {
      padding: 24px 16px;
      p {
        margin-bottom: 4px;
        font-size: 13px;
        line-height: 20px;
        color: #888888;
        word-break: keep-all;
        white-space: pre-line;
      }
      button {
        display: block;
        width: 90%;
        height: 48px;
        margin: 24px auto 0;
        font-size: 16px;
        line-height: 48px;
        border: none;
        border-radius: 5px;
      }
    }
  }
  .arrow180 {
      transform: rotate(180deg);
  }
}
.last-quiz {
  padding: 32px 0;
  margin-right: 5%;
  margin-left: 5%;
  width: 90%;
  .last-quiz-title {
    padding-bottom: 23px;
    font-size: 20px;
    font-weight: 600;
    color: #222;
  }
} 
</style>
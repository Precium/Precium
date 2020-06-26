<template>
<div class="quiz-index default-container">
  <div class="default-wrap">
    <go-to-last-page-nav :title = currentPageTitle />
    <div class="tips-and-notices">
      <div class="tips-and-notices-title">
        <img src="@/assets/digital-wallet.svg" alt="wallet">
        <p>도움말 및 유의사항</p>
      </div>
      <span class="divider-block" />
    </div>
    <div>
      <FaqList v-for="(contentList, b) in blockchainContent.faqBlockchain" :key="b"
      :bestFaq="contentList.title"
      :answer="contentList.content"
      :listId="b"
      />
    </div>
    <div class="notice">
      <p>유의사항</p>
      <p>• 지갑<br>
        - 지갑 생성까지 일정 시간이 소요될 수 있습니다.<br>
        - PCM지갑은 1개만 생성이 가능하며, 외부 지갑 주소는 3개까지 등록 가능합니다.<br>
        - PCM지갑은 오로지 PCM토큰만 입금할 수 있습니다.<br><br>

        • 입/출금<br>
        - PCM은 외부 거래소에서 거래 가능합니다.<br>
        - 외부 거래소에서 7renders 지갑 주소로 입금 후 입금확인 버튼을 클릭해야 PCM 수가 갱신되며, PCM 수는 소수점 제외하고 1의 자리만 카운트합니다.<br>
        - 출금 신청을 통해 다른 지갑 혹은 외부 거래소로 전송 가능합니다.<br>
      </p>
    </div>
  </div>
</div>
</template>

<script lang='js'>
import Vue from 'vue'
import Component from 'nuxt-class-component'
import {} from 'vuex-class'
import GoToLastPageNav from '@/components/GoToLastPageNav'
import FaqList from '@/components/FaqList';
import blockchainContent from '@/static/blockchainContent';

@Component({
  components: {
    GoToLastPageNav,
    FaqList,
  },
})
export default class HelpAboutWallet extends Vue {
  currentPageTitle = '지갑생성하기'
  blockchainContent = blockchainContent
  layout(context) {
    return 'onlyBody';
  }
}
</script>

<style lang='scss' scoped>
.quiz-index {
  height: 100vh;
}
.tips-and-notices {
  position: relative;
  width: 100%;
  .tips-and-notices-title {
    padding: 24px;
    display: flex;
    align-items: center;
    .icon {
      width: 40px;
      height: 40px;
      background-color: #FFDDA6;
    }
    p {
      padding-left: 12px;
      font-size: 20px;
      font-weight: bold;
      color: #222;
    }
  }
}
.notice {
  background-color: #FAFAFA;
  padding: 34px 16px;
  box-sizing: border-box;
  border-top: 1px solid #F2F2F2;
  p:first-child {
    padding-bottom: 4px;
    font-size: 14px;
    color: #D93F3F;
  }
  p:last-child {
    line-height: 18px;
    font-size: 13px;
    color: #888888;
  }
}
</style>

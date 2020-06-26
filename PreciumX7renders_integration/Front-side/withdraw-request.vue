<template>
<div class='withdraw-history-index default-container'>
  <div class="default-wrap">
    <go-to-last-page-nav :title = currentPageTitle />
    <withdraw-history v-for="request in requestHistory" :key="requestHistory.id" :request = request />
    <span class="divider-block" />
    <div class="notice cursor-pointer" @click="active = !active">
      <span class="notice-title">출금 요청 내역 안내</span>
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
import {} from 'vuex-class'
import GoToLastPageNav from '@/components/GoToLastPageNav'
import WithdrawHistory from '@/components/WithdrawHistory'

@Component({
  components: {
    GoToLastPageNav,
    WithdrawHistory
  },
})
export default class withdrawRequest extends Vue {
  currentPageTitle = '출금 요청 내역'
  active = false //vue-slide-up-down

  requestHistory = [
    {
      pcm: 2000,
      alias: '빗썸지갑',
      date: '2019.08.01 14:44:02'
    },
    {
      pcm: 1000,
      alias: '코인원지갑',
      date: '2019.08.01 14:44:02'
    },
    {
      pcm: 3000,
      alias: '프레시움지갑',
      date: '2019.08.01 14:44:02'
    }
  ]

  layout(context) {
    return 'onlyBody';
  }
}
</script>

<style lang='scss' scoped>
.withdraw-history-index {
  width: 100%;
  height: 100vh;
  .notice {
    position: relative;
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
}
.arrow180 {
  transform: rotate(180deg);
}
</style>

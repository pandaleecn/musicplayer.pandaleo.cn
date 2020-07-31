// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import axios from 'axios'
import VueAxios from 'vue-axios'
import router from './router'
import ElementUI from 'element-ui'
import utils from '@/utils'
import 'element-ui/lib/theme-chalk/index.css'

Vue.config.productionTip = false
Vue.use(ElementUI)
Vue.use(VueAxios, axios)


router.beforeEach(async (to, from, next) => {
  if (to.meta.requireAuth) {
    if (!utils.getCookie('token')) { // 没有登录则跳转/login页，进行登录
      console.log(123)
      next({
        name: 'Login',
        query: {
          redirect: to.name
        }
      })
    } else {
      if (store.state.UserProfile.length === 0) {
        await store.dispatch('getUserProfile')
      }
      next()
    }
  } else {
    next()
  }
})

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: {
    App
  },
  template: '<App/>'
})
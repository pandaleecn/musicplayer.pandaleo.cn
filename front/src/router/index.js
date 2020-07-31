import Vue from 'vue'
import Router from 'vue-router'
import Login from '@/components/Login'
import Home from '@/components/Home'
import Admin from '@/components/Admin'

import RoleMange from '@/components/RoleMange/RoleMange'
import UserMange from '@/components/UserMange/UserMange'
import SongMange from '@/components/SongMange/SongMange'
import SheetMange from '@/components/SheetMange/SheetMange'

const parentComponent = {
  template: `<router-view></router-view>`
};
Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '',
      component: Admin,
      children: [{
        path: '/',
        name: 'Home',
        meta: {
          titile: '主页',
          requireAuth: true,
          cid: 1
        },
        component: Home
      },
      {
        path: '/SongMange',
        name: 'SongMange',
        meta: {
          title: '歌曲管理',
          requireAuth: true,
          cid: 2
        },
        component: SongMange
      },
      {
        path: '/SheetMange',
        name: 'SheetMange',
        meta: {
          title: '歌单管理',
          requireAuth: true,
          cid: 2
        },
        component: SheetMange
      },
      {
        path: '/UserMange',
        name: 'UserMange',
        meta: {
          title: '用户管理',
          requireAuth: true,
          cid: 2
        },
        component: UserMange
      },
      {
        path: '/RoleMange',
        name: 'RoleMange',
        meta: {
          title: '角色管理',
          requireAuth: true,
          cid: 2
        },
        component: RoleMange
      }]
    },
    {
      path: '/Login',
      name: 'Login',
      meta: {
        title: '登录页',
        requireAuth: true,
      },
      component: Login
    }
  ]
})
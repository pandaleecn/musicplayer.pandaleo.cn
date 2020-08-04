import Vue from "vue";
import Router from "vue-router";
import Login from "@/components/Login";
import Admin from "@/components/Admin";
import Home from "@/components/Home";

// 系统设置
import PermissionsMange from "@/components/SetUp/PermissionsMange";
import AddPermissions from "@/components/SetUp/PermissionsMange/AddPermissions";
import RoleMange from "@/components/SetUp/RoleMange";
import AddRole from "@/components/SetUp/RoleMange/AddRole";
import UsersMange from "@/components/SetUp/UsersMange";
import AddUsers from "@/components/SetUp/UsersMange/AddUsers";
import SongMange from "@/components/ResourceMange/SongMange";
import AddSong from "@/components/ResourceMange/SongMange/AddSong";
import PlaylistMange from "@/components/ResourceMange/PlaylistMange";
import AddPlaylist from "@/components/ResourceMange/PlaylistMange/AddPlaylist";
import AlbumMange from "@/components/ResourceMange/AlbumMange";
import AddAlbum from "@/components/ResourceMange/AlbumMange/AddAlbum";
import ArtistMange from "@/components/ResourceMange/ArtistMange";
import AddArtist from "@/components/ResourceMange/ArtistMange/AddArtist";
import LyricMange from "@/components/ResourceMange/LyricMange";
import AddLyric from "@/components/ResourceMange/LyricMange/AddLyric";

const parentComponent = {
  template: `<router-view></router-view>`
};
Vue.use(Router);

export default new Router({
  // mode: 'history',
  routes: [
    {
      path: "",
      component: Admin,
      children: [
        {
          path: "/",
          name: "Home",
          meta: {
            title: "主页",
            requireAuth: true,
            cid: 1
          },
          component: Home
        },
        {
          path: "ResourceMange",
          name: "ResourceMange",
          meta: {
            title: "资源管理",
            requireAuth: true
          },
          component: parentComponent,
          children: [
            {
              path: "SongMange",
              name: "SongMange",
              meta: {
                title: "歌曲管理",
                requireAuth: true,
                cid: 11
              },
              component: SongMange
            },
            {
              path: "AddSong",
              name: "AddSong",
              meta: {
                title: "新增歌曲",
                requireAuth: true,
                cid: 11
              },
              component: AddSong
            },
            {
              path: "EditSong",
              name: "EditSong",
              meta: {
                title: "编辑歌曲",
                requireAuth: true,
                cid: 11
              },
              component: AddSong
            },
            {
              path: "PlaylistMange",
              name: "PlaylistMange",
              meta: {
                title: "歌单管理",
                requireAuth: true,
                cid: 12
              },
              component: PlaylistMange
            },
            {
              path: "AddPlaylist",
              name: "AddPlaylist",
              meta: {
                title: "新增歌单",
                requireAuth: true,
                cid: 12
              },
              component: AddPlaylist
            },
            {
              path: "EditPlaylist",
              name: "EditPlaylist",
              meta: {
                title: "编辑歌单",
                requireAuth: true,
                cid: 12
              },
              component: AddPlaylist
            },
            ,
            {
              path: "AlbumMange",
              name: "AlbumMange",
              meta: {
                title: "专辑管理",
                requireAuth: true,
                cid: 13
              },
              component: AlbumMange
            },
            {
              path: "AddAlbum",
              name: "AddAlbum",
              meta: {
                title: "新增专辑",
                requireAuth: true,
                cid: 13
              },
              component: AddAlbum
            },
            {
              path: "EditAlbum",
              name: "EditAlbum",
              meta: {
                title: "编辑歌单",
                requireAuth: true,
                cid: 13
              },
              component: AddAlbum
            },
            {
              path: "ArtistMange",
              name: "ArtistMange",
              meta: {
                title: "歌手管理",
                requireAuth: true,
                cid: 14
              },
              component: ArtistMange
            },
            {
              path: "AddArtist",
              name: "AddArtist",
              meta: {
                title: "新增歌单",
                requireAuth: true,
                cid: 14
              },
              component: AddArtist
            },
            {
              path: "EditArtist",
              name: "EditArtist",
              meta: {
                title: "编辑歌手",
                requireAuth: true,
                cid: 14
              },
              component: AddArtist
            },
            {
              path: "LyricMange",
              name: "LyricMange",
              meta: {
                title: "歌词管理",
                requireAuth: true,
                cid: 15
              },
              component: LyricMange
            },
            {
              path: "AddLyric",
              name: "AddLyric",
              meta: {
                title: "新增歌词",
                requireAuth: true,
                cid: 15
              },
              component: AddLyric
            },
            {
              path: "EditLyric",
              name: "EditLyric",
              meta: {
                title: "编辑歌词",
                requireAuth: true,
                cid: 15
              },
              component: AddLyric
            }
          ]
        },
        {
          path: "SetUp",
          name: "SetUp",
          meta: {
            title: "设置",
            requireAuth: true
          },
          component: parentComponent,
          children: [
            {
              path: "PermissionsMange",
              name: "PermissionsMange",
              meta: {
                title: "权限管理",
                requireAuth: true,
                cid: 7
              },
              component: PermissionsMange
            },
            {
              path: "AddPermissions",
              name: "AddPermissions",
              meta: {
                title: "新建权限",
                requireAuth: true,
                cid: 7
              },
              component: AddPermissions
            },
            {
              path: "EditPermissions/:id",
              name: "EditPermissions",
              meta: {
                title: "编辑权限",
                requireAuth: true,
                cid: 7
              },
              component: AddPermissions
            },
            {
              path: "RoleMange",
              name: "RoleMange",
              meta: {
                title: "角色管理",
                requireAuth: true,
                cid: 8
              },
              component: RoleMange
            },
            {
              path: "AddRole",
              name: "AddRole",
              meta: {
                title: "新建角色",
                requireAuth: true,
                cid: 8
              },
              component: AddRole
            },
            {
              path: "EditRole/:id",
              name: "EditRole",
              meta: {
                title: "编辑角色",
                requireAuth: true,
                cid: 8
              },
              component: AddRole
            },
            {
              path: "UsersMange",
              name: "UsersMange",
              meta: {
                title: "账号管理",
                requireAuth: true,
                cid: 9
              },
              component: UsersMange
            },
            {
              path: "AddUsers",
              name: "AddUsers",
              meta: {
                title: "新建账号",
                requireAuth: true,
                cid: 9
              },
              component: AddUsers
            },
            {
              path: "EditUsers/:id",
              name: "EditUsers",
              meta: {
                title: "编辑账号",
                requireAuth: true,
                cid: 9
              },
              component: AddUsers
            }
          ]
        }
      ]
    },
    {
      path: "/login",
      name: "Login",
      meta: {
        title: "登录页",
        requireAuth: false
      },
      component: Login
    }
  ]
});

import Vue from "vue";
import Vuex from "vuex";
import utils from "@/utils";

Vue.use(Vuex);

const store = new Vuex.Store({
  state: {
    UserProfile: [],
    PermissionsData: {
      ListData: [],
      queryData: {},
      total: 0
    },
    RolesData: {
      ListData: [],
      queryData: {},
      total: 0
    },
    SongsData: {
      ListData: [],
      queryData: {},
      total: 0
    },
    AdminsData: {
      ListData: [],
      queryData: {},
      total: 0
    },
    PlaylistsData: {
      ListData: [],
      queryData: {},
      total: 0
    },
    AlbumsData: {
      ListData: [],
      queryData: {},
      total: 0
    },
    ArtistsData: {
      ListData: [],
      queryData: {},
      total: 0
    },
    QiniuToken: []
  },
  mutations: {
    QiniuToken(state, data) {
      state.QiniuToken = data;
    },
    UserProfile(state, data) {
      state.UserProfile = data;
    },
    PermissionsData(state, data) {
      state.PermissionsData.ListData = data;
    },
    RolesData(state, data) {
      state.RolesData.ListData = data;
    },
    SongsData(state, data) {
      state.SongsData.ListData = data;
    },
    AdminsData(state, data) {
      state.AdminsData.ListData = data;
    },
    PlaylistsData(state, data) {
      state.PlaylistsData.ListData = data;
    },
    AlbumsData(state, data) {
      state.AlbumsData.ListData = data;
    },
    ArtistsData(state, data) {
      state.ArtistsData.ListData = data;
    }
  },
  actions: {
    async getUserProfile({ state, commit }) {
      const data = await utils.getUserProfile();
      commit("UserProfile", data.data.data);
    },
    async getPermissions({ state, commit }, datas) {
      const data = await utils.getPermissions(datas);
      commit("PermissionsData", data.data.data);
    },
    async getRoles({ state, commit }, datas) {
      const data = await utils.getRoles(datas);
      commit("RolesData", data.data.data);
    },
    async getSongs({ state, commit }, datas) {
      const data = await utils.getSongs(datas);
      commit("SongsData", data.data.data);
    },
    async getAdmins({ state, commit }, datas) {
      const data = await utils.getAdmins(datas);
      commit("AdminsData", data.data.data);
    },
    async getPlaylists({ state, commit }, datas) {
      const data = await utils.getPlaylists(datas);
      commit("PlaylistsData", data.data.data);
    },
    async getAlbums({ state, commit }, datas) {
      const data = await utils.getAlbums(datas);
      commit("AlbumsData", data.data.data);
    },
    async getArtists({ state, commit }, datas) {
      const data = await utils.getArtists(datas);
      commit("ArtistsData", data.data.data);
    }
  }
});

export default store;

<template>
  <div class="EnterClass">
    <h1>{{ msg }}</h1>
    <el-button type="primary" @click="login">Show Channels</el-button>
    <el-table :data="channel_data" border justify="center" style="width: 100%">
      <el-table-column prop="user_name" label="ChannelName"></el-table-column>
      <el-table-column prop="pass_word" label="Description"></el-table-column>
      <el-table-column prop="avatar" label="Visible"></el-table-column>
    </el-table>
  </div>
</template>

<script>
import Vue from 'vue'
export default {
  name: 'EnterName',
  data () {
    return {
      msg: '集成element',
      checked: true,
      channel_data: []
    }
  },
  methods: {
    login: function () {
      var url = '/api/user?offset=0&limit=30'
      Vue.axios
        .get(url)
        .then((response) => {
          this.channel_data = response.data
          console.log('response: ' + response.data[0].title_chinese)
          console.log('response: ' + JSON.stringify(response))
        })
        .catch((error) => {
          console.log('error!' + error)
        })
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1,
h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>

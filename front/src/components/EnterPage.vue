<template>
  <div class="EnterClass">
    <h1>{{ msg }}</h1>
    <el-button type="primary" @click="login">Show Channels</el-button>
    <el-table :data="channel_data" border justify="center" style="width: 100%">
      <el-table-column prop="title_chinese" label="ChannelName"></el-table-column>
      <el-table-column prop="short_description_chinese" label="Description"></el-table-column>
      <el-table-column prop="visible" label="Visible"></el-table-column>
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
      // this.channel_data.push({
      //   title_chinese: 1,
      //   short_description_chinese: 0,
      //   visible: 3
      // })
      var url = '/api/category/test'
      Vue.axios
        .get(url)
        .then((response) => {
          this.channel_data = response.data.results
          console.log('response: ' + response.data.results)
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

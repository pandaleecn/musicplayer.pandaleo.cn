<template>
  <div>
    <div class="page-title">
      <i class="iconfont icon-arrow-left cursor" @click="$router.go(-1)" /> {{ $route.meta.title }}
    </div>
    <div class="know-container" v-loading="loading">
      <div class="auditDaka">
        <el-form @submit.native.prevent class="margin-top-lg" :label-position="labelPosition" label-width="130px" :model="ruleForm"
          ref="ruleForm" :rules="rules">

          <el-form-item prop="Name" class="margin-left-lg margin-top must" label="歌曲名">
            <el-input class="form-input-h" v-model="ruleForm.Name" placeholder="请输入" />
            <p class="form_p_g">请输入歌曲名称！</p>
          </el-form-item>

          <el-form-item prop="Url" class="margin-left-lg margin-top must" label="歌曲链接">
            <el-input class="form-input-h" v-model="ruleForm.Url" placeholder="请输入" />
            <p class="audio-box">
              <audio ref="audio" controls />
            </p>
            <p class="form_p_g">请输入歌曲链接或直接上传！</p>
            <el-upload class="upload-pic" :action="uploaddomain" :data="QiniuData" :on-remove="handleRemove" :on-error="uploadError"
              :on-success="uploadSuccess" :before-remove="beforeRemove" :before-upload="beforeSongUpload" :limit="3" multiple
              :on-exceed="handleExceed" :file-list="fileList">
              <el-button size="small" type="primary">上传歌曲</el-button>
            </el-upload>
          </el-form-item>

          <el-form-item prop="" class="margin-left-lg margin-top rolement" label="选择歌手">
            <!-- <div class="tree-box"> -->
            <!--  :default-expanded-keys="[2, 3]" -->
            <!-- 默认展开的的节点 -->
            <!-- <el-tree ref="tree" :data="data" show-checkbox node-key="Id" :props="defaultProps" :default-checked-keys="per_ids"> -->
            <!-- </el-tree> -->
            <!-- </div> -->
            <div>
              <el-select v-model="ruleForm.ArtistID" placeholder="请选择">
                <el-option v-for="item in ArtistsData.ListData" :key="item.Id" :label="item.Name" :value="item.Id">
                </el-option>
              </el-select>
            </div>
          </el-form-item>

          <el-form-item class="text-center margin-top-lg">
            <el-button :disabled="loading" type="primary" @click="submitForm('ruleForm')">提交</el-button>
            <el-button @click="resetForm('ruleForm')">返回</el-button>
          </el-form-item>
        </el-form>
      </div>
    </div>
  </div>
</template>

<script>
import { mapActions, mapState } from 'vuex'
import utils from '@/utils'
import * as qiniu from 'qiniu-js'

export default {
  data () {
    return {
      data: [],
      labelPosition: 'right',
      per_ids: [],
      ruleForm: {
        Name: '',
        Url: '',
        ID: this.$route.params.id,
        ArtistID: ''
      },
      rules: {
        Name: [{
          required: true,
          message: '请输入名',
          trigger: 'blur'
        }],
        Url: [{
          message: '请输入标识',
          trigger: 'blur'
        }],
        Artist: [{
          trigger: 'blur'
        }]
      },
      loading: false,
      resetdata: false,
      defaultProps: {
        children: 'children',
        label: 'Description'
      },
      QiniuData: {
        key: '', // 图片名字处理
        token: '' // 七牛云token
      },
      uploaddomain: 'https://up-z0.qiniup.com', // 七牛云的上传地址（华南区）
      qiniuaddr: 'http://files.pandaleo.cn', // 七牛云的图片外链地址
      uploadPicUrl: '', // 提交到后台图片地址
      originName: '',
      fileList: []
    }
  },
  computed: {
    ...mapState([
      'ArtistsData'
    ])
  },
  methods: {
    ...mapActions([
      'getArtists'
    ]),
    async getData () {
      if (this.$route.params.id) {
        this.loading = true
        const data = await utils.getSongsDetail(this.$route.params.id)
        let ruleForm = data.data.data
        this.ruleForm = ruleForm
        this.$refs.audio.src = data.data.data.Url

        if (ruleForm.Perms) {
          ruleForm.Perms.forEach(e => {
            this.per_ids.push(e.Id)
          })
        }
      } else {
        this.data.forEach(e => {
          this.per_ids.push(e.Id)
        })
      }

      this.loading = false
    },
    // 提交表单
    submitForm (formName) {
      let id = this.$route.params.id
      this.$refs[formName].validate((valid) => {
        if (valid) {
          if (id) {
            this.putSongs()
          } else {
            this.postSongs()
          }
        } else {
          console.log('error submit!!')
          return false
        }
      })
    },
    async postSongs () {
      this.loading = true
      const data = await utils.postSongs(this.ruleForm)
      if (data.data.status) {
        this.$message({
          type: 'success',
          message: data.data.msg
        })
        this.$router.push({
          name: 'SongMange'
        })
      } else {
        this.$message({
          type: 'info',
          message: data.data.msg
        })
      }
      this.loading = false
    },
    async putSongs () {
      this.loading = true
      const data = await utils.putSongs(this.ruleForm)
      if (data.data.status) {
        this.$message({
          type: 'success',
          message: data.data.msg
        })
        await this.$router.push({
          name: 'SongMange'
        })
      } else {
        this.$message({
          type: 'info',
          message: data.data.msg
        })
      }
      this.loading = false
    },
    async getQiniuToken () {
      this.loading = true
      const data = await utils.getQiniuToken()
      if (data.data.status) {
        this.QiniuData.token = data.data.data
      } else {
        console.log('获取七牛 token 失败！')
      }

      this.loading = false
    },
    handleRemove (file, fileList) {
      this.uploadPicUrl = ''
    },
    handleExceed (files, fileList) {

    },
    beforeSongUpload (file) {
      var date = new Date().format('yyyy-MM-dd')
      var time = new Date().toLocaleTimeString().replace(/:/g, '-')
      let dot = file.name.lastIndexOf('.')
      let fileOriginName = file.name.substring(0, dot)
      this.originName = fileOriginName
      var fileName = fileOriginName + '-' + date + time + '.mp3'
      this.QiniuData.key = fileName
    },
    uploadSuccess (response, file, fileList) {
      this.ruleForm.Url = `${this.qiniuaddr}/${this.QiniuData.key}`
      this.ruleForm.Name = this.originName
    },
    uploadError (err, file, fileList) {
      console.log('歌曲上传失败！')
    },
    beforeRemove (file, fileList) {
      // return this.$confirm(`确定移除 ${ file.name }？`);
    },
    // 提交数据到后台
    handleSubmit () {

    },
    resetForm (formName) {
      // this.$refs[formName].resetFields();
      if (this.resetdata) {
        this.$confirm('真的要退出此次编辑？', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消'
        }).then(() => {
          this.$router.go(-1)
        }).catch(() => {
          this.$message({
            type: 'info',
            message: '已取消'
          })
        })
      } else {
        this.$router.go(-1)
      }
    },
    uploadFile () {

    }
  },
  mounted: function () {
    this.getArtists()
    this.getData()
    this.getQiniuToken()
  },
  watch: {
    ruleForm: {
      handler (curVal, oldVal) {
        this.resetdata = true
      },
      deep: true
    }
  }
}
</script>
<style lang="less" scoped>
.auditDaka {
  &-arrow {
    & .icon-07jiantouxiangzuo {
      vertical-align: middle;
      font-size: 24px;
      color: #444444;
    }
  }
}
</style>
<style scoped>
.form-input-h {
  width: 400px;
}

.form_p_g {
  font-size: 14px;
  color: #888;
  clear: both;
}

.tree-box {
  margin-top: 10px;
  border: 1px solid #e3e3e3;
  border-radius: 4px;
  width: 800px;
  padding: 10px;
  min-height: 220px;
  max-height: 500px;
  overflow: scroll;
}

.audio-box {
  margin-top: 10px;
  width: 800px;
}

.avatar-uploader {
  border: 1px dashed #d9d9d9;
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
  display: inline-block;
  line-height: 1;
}

.avatar-uploader:hover {
  border-color: #409eff;
}

.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 100px;
  height: 100px;
  line-height: 100px;
  text-align: center;
}

.avatar {
  width: 100px;
  height: 100px;
  display: block;
}

.colrecom_form {
  width: 200px;
  border: 1px solid #ccc;
  line-height: initial;
  display: inline-block;
  position: relative;
  margin: 0 15px 1rem 0;
  float: left;
  height: 141px;
  border-radius: 4px;
}

.colrecom_form_add {
  width: 200px;
  border: 1px dotted #ccc;
  line-height: initial;
  display: inline-block;
  position: relative;
  margin: 0 15px 1rem 0;
  float: left;
  height: 141px;
  text-align: center;
  border-radius: 4px;
}

.colrecom_form_add > .el-icon-plus {
  height: 80px;
}

.colrecom_form > img {
  width: 100%;
}

.colrecom_form > p {
  line-height: normal;
  text-align: center;
  padding: 0 3px;
  word-wrap: normal;
  text-overflow: ellipsis;
  white-space: nowrap;
  overflow: hidden;
  text-align: justify;
}

.colrecom_form > .el-icon-error {
  color: #ff7043;
  position: absolute;
  top: -7px;
  right: -7px;
  font-size: 18px;
}

.input_search_video {
  width: 50%;
}

.span_search_video {
  color: #888;
  margin-left: 5%;
}

.select_show_video {
  width: 685px;
  margin: 0 auto;
}

.show_page {
  margin-top: 2rem;
  text-align: center;
}
</style>

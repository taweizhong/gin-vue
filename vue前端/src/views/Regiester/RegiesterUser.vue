<script>

export default {
  data() {
    return {
      user: {
        name: '',
        iphone: '',
        password: '',
      },
      show: false
    }
  },
  methods: {
    onSubmit() {
      if (this.user.iphone.length !== 11) {
        this.show = true
        return
      }
      this.show = false
      alert("密码正确")
      const api = "http://127.0.0.1:1016/register"
      this.axios.post(api, { ...this.user }).then(res => {
        console.log(res.data)
      }).catch(err => {
        this.$attrs.toast(err.response.data.msg, {
          title: "数据验证失败",
          variant: 'danger',
          solid: true
        })
      });
    },

  }
}
</script>

<template>
  <div class="container card mt-2" style="width: 40rem;">
    <div class="card-body">
      <h5 class="card-title">注册</h5>
      <div class="container col-5 align-items-cente">
        <form>
          <div class="mb-3">
            <label for="exampleInputEmail1" class="form-label">用户名</label>
            <input type="text" class="form-control" id="exampleInputEmail1" aria-describedby="emailHelp"
              v-model="user.name">
          </div>
          <div class="mb-3">
            <label for="exampleInputEmail1" class="form-label">电话号码</label>
            <input type="text" class="form-control" id="exampleInputEmail1" aria-describedby="emailHelp"
              v-model="user.iphone">
          </div>
          <div v-if="show" class="alert alert-primary" role="alert">
            密码错误
          </div>
          <div class="mb-3 ">
            <label for="exampleInputPassword1" class="form-label">密码</label>
            <input type="password" class="form-control" id="exampleInputPassword1" v-model="user.password">
          </div>
          <div class="d-grid gap-2">
            <button class="btn btn-primary" type="button" @click="onSubmit">注册</button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<style>

</style>
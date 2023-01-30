<script>
export default {
    data() {
        return {
            user: {
                name: '',
                iphone: '',
                password: '',
            },
            show: true
        }

    },
    methods: {
        onSubmit(evt) {
            evt.preventDefault()
            console.log(JSON.stringify(this.user))
        },
        onReset(evt) {
            evt.preventDefault()
            // Reset our form values
            this.user.name = ''
            this.user.iphone = ''
            this.user.password = ''
            // Trick to reset/clear native browser form validation state
            this.show = false
            this.$nextTick(() => {
                this.show = true
            })
        },
        password() {

        }
    }
}


</script>

<template>
    <div class="container card mt-2" style="width: 40rem;">
        <div class="card-body">
            <h5 class="card-title">登录</h5>
            <div class="container col-5 align-items-cente">
                <form @submit="onSubmit" @reset="onReset" v-if="show">
                    <div class="mb-3">
                        <label for="exampleInputEmail1" class="form-label">用户名</label>
                        <input type="text" class="form-control" id="exampleInputEmail1" aria-describedby="emailHelp"
                            v-model="user.name">
                    </div>
                    <div class="mb-3">
                        <label for="exampleInputEmail1" class="form-label">电话号码</label>
                        <input type="number" class="form-control" id="exampleInputEmail1" aria-describedby="emailHelp"
                            v-model="user.iphone">
                    </div>
                    <div class="mb-3 ">
                        <label for="exampleInputPassword1" class="form-label">密码</label>
                        <input type="password" class="form-control" id="exampleInputPassword1" v-model="user.password">
                    </div>
                    <div class="mb-3 form-check ">
                        <input type="checkbox" class="form-check-input" id="exampleCheck1" @click="password">
                        <label class="form-check-label" for="exampleCheck1">查看</label>
                    </div>
                    <div class="d-grid gap-2">
                        <button class="btn btn-primary" type="button">登录</button>
                    </div>
                </form>
            </div>
        </div>
    </div>
</template>

<style>

</style>
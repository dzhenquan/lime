<template>
    <div class="personal" v-if="info">
        <mt-header title="个人中心">
            <mt-button icon="back" slot="left" @click="goBack()"></mt-button>
        </mt-header>
        <div class="info">
            <div class="head">
                <img :src="info.faceicon" alt="">
            </div>
            <div class="name">{{info.username}}</div>
            <div class="item">
                <mt-cell title="我的书架" is-link to="/shelf"></mt-cell>
            </div>
        </div>
        <mt-button class="btn" type="danger" @click="loginOut()">退出登录</mt-button>
    </div>
</template>

<script>
    import {MessageBox} from 'mint-ui'
    import {userInfo} from "../api";
    import { removeToken } from '@/utils/auth'

    export default {
        name: "Personal",
        data() {
            return {
                info: null
            }
        },
        methods: {
            
        },
        created() {
            this.getUserInfo()
        },
        methods: {
            async getUserInfo() {
                this.loading = true;
                try {
                  const list = await userInfo();
                  this.info = list.data.result;
                } finally {
                  this.loading = false;
                }
            },
            goBack() {
                this.$router.go(-1);
            },
            loginOut() {
                MessageBox({
                    title: '提示',
                    message: '确定执行此操作?',
                    showCancelButton: true
                }).then(action => {
                    if (action == 'confirm') {
                        removeToken();
                        this.$router.push('/login');
                    } else {
                        return false
                    }
                })
            }
        }
    }
</script>

<style lang="less">
    .personal {
        .info {
            padding-top: 80px;

            .head {
                margin: 0 auto;
                width: 80px;
                height: 80px;
                -webkit-border-radius: 50%;
                -moz-border-radius: 50%;
                border-radius: 50%;
                overflow: hidden;

                img {
                    width: 100%;
                    height: 100%;
                }
            }

            .name {
                text-align: center;
                font-size: 20px;
                line-height: 1.8;
                margin-bottom: 60px;
            }

            .item {
                .mint-cell {
                    border-bottom: 1px solid #d9d9d9;

                    .mint-cell-wrapper {
                        background: none;
                    }
                }
            }
        }

        .btn {
            display: block;
            width: 80%;
            margin: 40px auto 0;
        }
    }
</style>

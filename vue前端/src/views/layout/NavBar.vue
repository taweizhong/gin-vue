<template>
  <div class="container text-center">
    <ul class="nav justify-content-center">
      <li class="nav-item">
        <a class="nav-link" @click="$router.replace({ name: 'login' })">主页</a>
      </li>
      <li class="nav-item">
        <a class="nav-link" @click="$router.replace({ name: 'login' })">登录</a>
      </li>
      <li class="nav-item">
        <a class="nav-link" @click="$router.replace({ name: 'regiester' })">注册</a>
      </li>
    </ul>
  </div>
</template>

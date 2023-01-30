import { createApp } from 'vue'
import App from './App.vue'

import router from './router'
import VuelidatePlugin from "@vuelidate/core";


import '/src/assets/styles.scss'
import "@popperjs/core";
import "bootstrap";

import axios from "axios";
import VueAxios from "vue-axios";


const app = createApp(App)

app.use(router)
app.use(VuelidatePlugin)
app.use(VueAxios, axios)


app.mount('#app')

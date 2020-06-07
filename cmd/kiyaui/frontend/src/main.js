import 'core-js/stable';
import 'regenerator-runtime/runtime';
import '@mdi/font/css/materialdesignicons.css';
import Vue from 'vue';
import Vuetify from 'vuetify';
import 'vuetify/dist/vuetify.min.css';
import store from "./store";


Vue.use(Vuetify);

import App from './App.vue';

Vue.config.productionTip = false;
Vue.config.devtools = true;

import Wails from '@wailsapp/runtime';

Wails.Init(() => {
	new Vue({
		store,
		vuetify: new Vuetify({
			icons: {
				iconfont: 'mdi'
			},
			theme: {
				dark: true
			}
		}),
		render: h => h(App)
	}).$mount('#app');
});

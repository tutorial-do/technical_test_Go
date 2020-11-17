import Vue from 'vue';
import Vuex from 'vuex';
import axios from 'axios';

Vue.use(Vuex);
// const baseURL = 'http://localhost:3000/';

export default new Vuex.Store({
  state: {
    allBuyers: [],
  },
  mutations: {
    SET_ALL_BUYERS(state, buyers) {
      state.allBuyers = buyers;
      console.log('HOla MAucli');
      console.log(state.allBuyers);
      console.log('Chao MAucli');
    },
  },
  actions: {
    async fetchAllBuyers(context) {
      try {
        const response = await axios.get('http://localhost:3000/allbuyers');
        const buyers = response.data.allBuyers;
        context.commit('SET_ALL_BUYERS', buyers);
      } catch (error) {
        console.log(error);
      }
    },
  },
  // getters: {
  //   getAllBuyers(state) {
  //     return state.allBuyers;
  //   },
  // },
});

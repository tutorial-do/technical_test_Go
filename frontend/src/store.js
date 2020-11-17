import Vue from 'vue';
import Vuex from 'vuex';
import axios from 'axios';

Vue.use(Vuex);
// const baseURL = 'http://localhost:3000/';

export default new Vuex.Store({
  state: {
    allBuyers: [],
    allProducts: [],
    allTransactions: [],
  },
  mutations: {
    SET_ALL_BUYERS(state, buyers) {
      state.allBuyers = buyers;
      console.log(state.allBuyers);
    },
    SET_ALL_PRODUCTS(state, products) {
      state.allProducts = products;
      console.log(state.allProducts);
    },
    SET_ALL_TRANSACTIONS(state, transactions) {
      state.allTransactions = transactions;
      console.log(state.allTransactions);
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
    async fetchAllProducts(context) {
      try {
        const response = await axios.get('http://localhost:3000/allproducts');
        const products = response.data.allProducts;
        context.commit('SET_ALL_PRODUCTS', products);
      } catch (error) {
        console.log(error);
      }
    },
    async fetchAllTransactions(context) {
      try {
        const response = await axios.get('http://localhost:3000/alltransactions');
        const transactions = response.data.allTransactions;
        context.commit('SET_ALL_TRANSACTIONS', transactions);
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

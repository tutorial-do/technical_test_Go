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
    activeBuyer: {
      name: '',
      id: '',
      age: '',
    },
    transactionsByBuyerID: [],
    productsByTransaction: [],
    buyersIP: [],
    currentIP: '',
  },
  mutations: {
    SET_ALL_BUYERS(state, buyers) {
      state.allBuyers = buyers;
    },
    SET_ALL_PRODUCTS(state, products) {
      state.allProducts = products;
    },
    SET_ALL_TRANSACTIONS(state, transactions) {
      state.allTransactions = transactions;
    },
    SET_CURRENT_BUYER: (state, payload) => {
      state.activeBuyer.id = payload.id;
      state.activeBuyer.name = payload.name;
      state.activeBuyer.age = payload.age;
      console.log(state.activeBuyer);
    },
    SET_PRODUCTS_BY_TRANSACTION: (state, currentProducts) => {
      state.productsByTransaction = currentProducts;
      console.log(state.productsByTransaction);
    },
    SET_TRANSACTIONS_BY_BUYER: (state, transactionsByBuyerID) => {
      state.transactionsByBuyerID = transactionsByBuyerID;
    },
    SET_BUYERS_BY_IP: (state, buyersByIP) => {
      state.buyersIP = buyersByIP;
    },
    SET_CURRENT_IP: (state, currentIPAddress) => {
      state.currentIP = currentIPAddress;
    },
  },
  actions: {
    setCurrentBuyer(context, payload) {
      const currentBuyerID = payload.id;
      context.dispatch('getTransactionsByBuyerID', currentBuyerID);
      context.commit('SET_CURRENT_BUYER', payload);
    },
    getTransactionsByBuyerID(context, currentBuyerID) {
      context.dispatch('fetchAllTransactions');
      const transact = context.state.allTransactions;
      const transactionsByBuyerID = transact.filter((trans) => trans.buyerID === currentBuyerID);
      context.commit('SET_TRANSACTIONS_BY_BUYER', transactionsByBuyerID);
    },
    getBuyersByIP(context, ipAddress) {
      const transactions = context.state.allTransactions;
      const buyersByIP = transactions.filter((trans) => trans.ip === ipAddress);
      const buyers = context.state.allBuyers;
      const buyersFinal = [];
      for (let i = 0; i < buyersByIP.length; i += 1) {
        for (let j = 0; j < buyers.length; j += 1) {
          if (buyers[j].id === buyersByIP[i].buyerID) {
            buyersFinal.push(buyers[j]);
          }
        }
      }
      context.commit('SET_CURRENT_IP', ipAddress);
      context.commit('SET_BUYERS_BY_IP', buyersFinal);
    },
    getProductsByTransaction(context, currentProductsIds) {
      const productsIDs = currentProductsIds;
      const allProd = context.state.allProducts;
      const currentProducts = allProd.filter((p) => productsIDs.includes(p.id) === true);
      context.commit('SET_PRODUCTS_BY_TRANSACTION', currentProducts);
    },
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

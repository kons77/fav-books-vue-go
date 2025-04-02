<template>
  <Header />

  <div>
    <router-view @success="success" @error="error" @warning="warning" />
  </div>
  <Footer />
  
</template>

<script>
  import Header from "./components/Header.vue";
  import Footer from "./components/Footer.vue";
  import { store } from './components/store.js'
  import notie from 'notie'

  const getCookie = (name) => {
    return document.cookie.split("; ").reduce((r, v) => {
      const parts = v.split("=");
      return parts[0] === name ? decodeURIComponent(parts[1]) : r;
    }, "");
  }

  /*  const getCookie = (name) => {
        return Object.fromEntries(document.cookie.split("; ").map(c => c.split("=")))[name] || "";
      }; 
  */

  export default {
    name: 'App',
    components: {
      Header,
      Footer,
    },
    data() {
      return {
        store // Это не копия! Это ссылка на общий `store`
      }
    }, 
    beforeMount() {
      // check for a cookie
      let data = getCookie("_site_data");

      if (data !== "") {
        let cookieData = JSON.parse(data);

        // update store
        store.token = cookieData.token.token;
        store.user = {
          id: cookieData.user.id,
          first_name: cookieData.user.first_name,
          last_name: cookieData.user.last_name,
          email: cookieData.user.email,
        }
      }
    },
    mounted() {
      //fooBarTest();    
    },
    methods: {
      success(msg){
        notie.alert({
          type: 'success',
          text: msg,
        })
      },
      error(msg) {
        notie.alert({
          type: 'error',
          text: msg,
        })
      },
      warning(msg) {
        notie.alert({
          type: 'warning',
          text: msg,
        })
      },
    },
  }

  function fooBarTest() {
    const payload = {
          foo: "bar",
        }

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("Authorization", "Bearer " + store.token)

    const requestOptions = {
      method: "POST", 
      body: JSON.stringify(payload),
      headers: headers
    }

    fetch(`${store.apiBaseURL}/admin/foo`, requestOptions)
    .then(response => response.json())
    .then((response) => {
      console.log(response);
    })
  } 
  
</script>

<style>

</style>


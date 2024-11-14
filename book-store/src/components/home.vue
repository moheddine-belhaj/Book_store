<template>
    <div>
      <h1>Welcome to Home Page</h1>
      <button @click="navigateToHelloWorld">Go to Hello World</button>
    </div>
  </template>
  
  <script>
  export default {
    name: 'Home',
    methods: {
      navigateToHelloWorld() {
        this.$router.push({ name: 'HelloWorld' });
      }
    }
  };
  </script>
  
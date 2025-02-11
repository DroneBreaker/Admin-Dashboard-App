<template>
    <div class="login pt-[13rem] mx-[35%]">
        <form action="POST" class="bg-orange-400 pt-4 w-[27rem] h-[18rem] text-center rounded-md">
            <h1 class="login-text text-2xl mb-4">Welcome to Login</h1>

            <div>
                <input class="w-[80%]" type="text" placeholder="Username">
            </div>

            <div>
                <input class="w-[80%]" type="text" placeholder="Password">
            </div>

            <p></p>

            <Button buttonText="Submit" color="white" textColor="black" />
        </form>
    </div>
  </template>
  
  <script setup>
  import Button from '../components/Button.vue'
  </script>
  
  <style lang="css" scoped> 
  input {
    margin-bottom: 20px;
    border-radius: 10px;
    height: 50px;
    padding-left: 20px;
    border: 2px solid white;
  }
  </style>
  
  
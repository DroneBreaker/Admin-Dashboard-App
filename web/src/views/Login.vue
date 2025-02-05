<template>
    <div>
        <div class="login">
            <form action="POST">
                <h1 class="login-text">Welcome to Login</h1>

                <div>
                    <input type="text" placeholder="Username">
                </div>

                <div>
                    <input type="text" placeholder="Password">
                </div>

                <p></p>

                <Button buttonText="Submit" />
            </form>
        </div>
    </div>
</template>

<script setup>
import Button from '../components/Button';
</script>

<style lang="css" scoped>
.login {
    padding-top: 180px;
    margin-left: 36%;
    justify-content: center;
    align-items: center;
    height: 220px;
    width: 360px;
}

.login-text {
    font-size: 24px;
}

form {
    background-color: cadetblue;
    padding-top: 20px;
    text-align: center;
    border-radius: 10px;
}

input {
    margin-bottom: 20px;
    border-radius: 10px;
    height: 40px;
    width: 250px;
    padding-left: 20px;
}
</style>
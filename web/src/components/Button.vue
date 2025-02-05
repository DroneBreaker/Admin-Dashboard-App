<template>
    <button :style="{ backgroundColor: color }" class="button">
        {{ buttonText }}
    </button>
</template>

<script setup>
// Define props
const props = defineProps({
    buttonText: {
        type: String,
        required: true
    }, 
    color: {
        type: String,
        default: 'white'
    }
})
</script>

<style lang="css" scoped>
button {
    padding: 10px;
    border: none;
    cursor: pointer;
    text-align: center;
    font-size: 15px;
    transition: background-color 0.3s ease;
    border-radius: 10px;
    margin-bottom: 10px;
    width: 120px;
    height: 45px;
}
</style>
<template>
    <button :style="{ backgroundColor: color, color: textColor }" class="px-2 p-4 w-[8rem] rounded-md mb-4">
        {{ buttonText }}
    </button>
</template>

<script setup>
const props = defineProps({
    buttonText: {
        type: String,
        required: true
    },
    color: {
        type: String,
        default: 'blue'
    },
    textColor: {
        type: String,
        default: 'white'
    }
})
</script>

<style lang="scss" scoped>

</style>
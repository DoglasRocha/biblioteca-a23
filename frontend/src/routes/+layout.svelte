<script lang="ts">

</script>

<div class="main-wrapper">
    <slot></slot>
</div>

<style>
    :global(*) {
        padding: 0;
        margin: 0;
    }

    :global(body) {
        background-color: #F6CE5F;
    }

    .main-wrapper {
        min-height: 100vh;
        min-width: 100vw;
        display: flex;
        justify-content:space-around;
        align-items: center;
        flex-direction: column;
    }
</style>
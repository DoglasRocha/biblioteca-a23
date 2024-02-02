<script>
	import Footer from '$lib/components/footer.svelte';
</script>

<div class="main-wrapper">
	<slot />
	<Footer />
</div>

<style>
	:global(*) {
		padding: 0;
		margin: 0;
	}

	:global(body) {
		background-color: #f6ce5f;
	}

	.main-wrapper {
		min-height: 100vh;
		min-width: 100vw;
		display: flex;
		justify-content: space-between;
		align-items: center;
		flex-direction: column;
	}
</style>

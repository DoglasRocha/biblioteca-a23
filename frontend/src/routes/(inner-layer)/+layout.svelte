<script>
	import Navbar from '../../lib/components/navbar.svelte';
</script>

<div class="main-wrapper">
	<Navbar />
	<slot />
	<Navbar />
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
		justify-content: space-around;
		align-items: center;
		flex-direction: column;
	}
</style>

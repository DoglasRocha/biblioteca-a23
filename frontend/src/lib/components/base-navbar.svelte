<script>
</script>

<nav class="navbar navbar-expand-lg w-100 bg-body-tertiary">
	<div class="container-fluid">
		<a href="/">
			<img class="navbar-brand logo" src="/logo-menor.png" alt="Logo do Cursinho Popular A23" />
		</a>
		<button
			class="navbar-toggler"
			type="button"
			data-bs-toggle="collapse"
			data-bs-target="#navbarSupportedContent"
			aria-controls="navbarSupportedContent"
			aria-expanded="false"
			aria-label="Toggle navigation"
		>
			<span class="navbar-toggler-icon"></span>
		</button>
		<div class="collapse navbar-collapse" id="navbarSupportedContent">
			<ul class="navbar-nav me-auto mb-2 mb-lg-0">
				<slot />
			</ul>
		</div>
	</div>
</nav>

<style>
	.logo {
		width: 75px;
	}
</style>

<script>
	export let baseLink = '';
</script>

<nav class="navbar navbar-expand-lg w-100 bg-body-tertiary">
	<div class="container-fluid">
		<a href={`${baseLink}/`}>
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
			<div class="p-2">
				<a class="d-flex align-items-center nav-link fw-bold" href={`${baseLink}/minhaconta`}>
					<i class="fa-solid fa-circle-user me-2 my-account-icon"></i>
					Minha Conta
				</a>
				<a href="/sair" class="nav-link text-danger text-decoration-underline text-end">Sair</a>
			</div>
		</div>
	</div>
</nav>

<style>
	.logo {
		width: 75px;
	}

	.my-account-icon {
		font-size: 20px;
	}
</style>

<script>
	import FormField from '$lib/components/form-field.svelte';
	//import Logo from '/logo.png';

	let email = '',
		senha = '';
</script>

<div class="card">
	<div class="card-body">
		<h1>Biblioteca do Cursinho A23</h1>

		<div class="d-flex justify-content-center mt-3">
			<img src="/logo.png" alt="Logo do Cursinho Popular Alicerce 23" class="logo" />
		</div>

		<div class="mt-5">
			<h3>Login</h3>
		</div>

		<FormField
			label="Endereço de Email"
			name="email"
			type="email"
			bind:value={email}
			placeholder="exemplo@email.com"
		/>

		<FormField label="Senha" name="senha" type="password" bind:value={senha} />

		<div class="buttons mt-3">
			<a class="btn btn-secondary" href="/cadastro"> Cadastrar </a>
			<button
				class="btn btn-primary"
				on:click={() => {
					console.log(email);
				}}>Login</button
			>
		</div>
	</div>
</div>

<style>
	.buttons {
		display: flex;
		justify-content: space-between;
	}

	.logo {
		width: 200px;
	}
</style>

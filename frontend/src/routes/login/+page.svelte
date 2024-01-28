<script>
	import FormField from '$lib/components/form-field.svelte';
	import Card from '../../lib/components/card.svelte';
	//import Logo from '/logo.png';

	let email = '',
		senha = '';
</script>

<Card>
	<h1 class="card-title">Biblioteca do Cursinho A23</h1>

	<div class="d-flex justify-content-center mt-3">
		<img src="/logo.png" alt="Logo do Cursinho Popular Alicerce 23" class="logo" />
	</div>

	<div class="mt-5">
		<h3>Login</h3>
	</div>

	<form>
		<FormField
			label="Endereço de Email"
			name="email"
			type="email"
			bind:value={email}
			placeholder="exemplo@email.com"
			required
		/>

		<FormField label="Senha" name="senha" type="password" bind:value={senha} required />
		<div class="buttons mt-3">
			<a class="btn btn-secondary" href="/cadastro"> Cadastrar </a>
			<button
				class="btn btn-primary"
				type="submit"
				on:click={() => {
					if (email && senha) {
						document.cookie = 'logged=true';
						document.location.href = '/';
					}
				}}>Login</button
			>
		</div>
	</form>
</Card>

<style>
	.buttons {
		display: flex;
		justify-content: space-between;
	}

	.logo {
		width: 200px;
	}
</style>

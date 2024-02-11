<script>
	import FormField from '$lib/components/form-field.svelte';
	import Card from '$lib/components/card.svelte';
	import BlankForm from '$lib/components/blank-form.svelte';
	import { api } from '$lib/utils/api';
	//import Logo from '/logo.png';

	let loginFields = {
		email: '',
		password: ''
	};

	let errorFromServer;
	const handleSubmit = async () => {
		let request;
		try {
			request = await api.post('/login', loginFields);

			if (request.status == 200) document.location.href = '/';
		} catch (error) {
			errorFromServer = error.response.data;
		}
	};
</script>

<Card>
	<h1 class="card-title">Biblioteca do Cursinho A23</h1>

	<div class="d-flex justify-content-center mt-3">
		<img src="/logo.png" alt="Logo do Cursinho Popular Alicerce 23" class="logo" />
	</div>

	<div class="mt-5">
		<h3>Login</h3>
	</div>

	<BlankForm>
		<FormField
			label="Endereço de Email"
			name="email"
			type="email"
			bind:value={loginFields.email}
			placeholder="exemplo@email.com"
			required
		/>

		<FormField
			label="Senha"
			name="senha"
			type="password"
			bind:value={loginFields.password}
			required
		/>
		{#if errorFromServer}
			<div class="mt-3 text-danger">
				<p>Ocorreu algum erro. Provavelmente isso ajude:</p>
				<p>{errorFromServer}</p>
			</div>
		{/if}
		<div class="buttons mt-3">
			<a class="btn btn-secondary" href="/cadastro"> Cadastrar </a>
			<button
				class="btn btn-primary"
				type="submit"
				on:click={() => {
					handleSubmit();
				}}>Login</button
			>
		</div>
	</BlankForm>
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

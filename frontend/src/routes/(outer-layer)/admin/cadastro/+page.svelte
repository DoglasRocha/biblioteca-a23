<script>
	import FormField from '$lib/components/form-field.svelte';
	import Card from '$lib/components/card.svelte';
	import BlankForm from '$lib/components/blank-form.svelte';
	import { api } from '$lib/utils/api.js';
	import { isValidUserInput } from '$lib/utils/helpers.js';

	let userData = {
		name: '',
		surname: '',
		email: '',
		password: ''
	};

	let isInvalid = {
		name: false,
		surname: false,
		email: false,
		password: false,
		passwordConfirmation: false
	};

	let passwordConfirmation, errorFromServer;

	const handleSubmit = async () => {
		if (!isValidUserInput(isInvalid)) return;

		if (userData.password != passwordConfirmation) {
			isInvalid.passwordConfirmation = true;
			return;
		}

		try {
			let request = await api.passwordConfirmation('/admin/cadastro', userData);

			if (request.status == 201) document.location.href = '/admin/login';
		} catch (error) {
			errorFromServer = error.response.data;
		}
	};
</script>

<Card>
	<h1>Cadastro</h1>
	<BlankForm>
		<div class="d-flex">
			<div class="me-1">
				<FormField
					name="first-name"
					label="Nome"
					placeholder="João"
					bind:value={userData.name}
					type="text"
					errorMsg="Seu nome deve possuir três ou mais letras"
					validation={/[\w ]{3,50}/}
					bind:isInvalid={isInvalid.name}
					required
				/>
			</div>
			<div class="ms-1">
				<FormField
					name="last-name"
					label="Sobrenome"
					placeholder="Silva"
					bind:value={userData.surname}
					type="text"
					errorMsg="Seu sobrenome deve possuir três ou mais letras"
					validation={/[\w ]{3,100}/}
					bind:isInvalid={isInvalid.surname}
					required
				/>
			</div>
		</div>
		<FormField
			name="email"
			label="Email"
			placeholder="exemplo@email.com"
			bind:value={userData.email}
			type="email"
			errorMsg="Email inválido"
			validation={/^[\w\-\.]+@([\w-]+\.)+[\w-]{2,}$/}
			bind:isInvalid={isInvalid.email}
			required
		/>
		<FormField
			name="password"
			label="Senha "
			bind:value={userData.password}
			type="password"
			errorMsg="Sua senha deve possuir mais de oito dígitos"
			validation={/.{8,}/}
			bind:isInvalid={isInvalid.password}
			required
		/>
		<FormField
			name="password-confirmation"
			label="Confirmação de senha"
			bind:value={passwordConfirmation}
			type="password"
			errorMsg="As senhas não são iguais"
			validation={/.*?/}
			bind:isInvalid={isInvalid.passwordConfirmation}
			required
		/>
		<div class="mt-3 d-flex justify-content-end">
			<button
				type="submit"
				class="btn btn-primary"
				on:click={() => {
					handleSubmit();
				}}>Cadastrar</button
			>
		</div>
	</BlankForm>
</Card>

<style>
</style>

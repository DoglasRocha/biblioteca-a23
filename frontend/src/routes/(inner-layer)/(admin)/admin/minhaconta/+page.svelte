<script>
	import FormField from '$lib/components/form-field.svelte';
	import Card from '$lib/components/card.svelte';
	import BlankForm from '$lib/components/blank-form.svelte';
	import { isValidUserInput } from '$lib/utils/helpers.js';
	import { api } from '$lib/utils/api.js';
	export let data;

	let userData = data.userData;

	let isInvalid = {
		name: false,
		surname: false,
		email: false,
		password: false,
		passwordConfirmation: false,
		newPassword: false
	};

	let newPassword = '',
		passwordConfirmation = '',
		errorFromServer;

	const handleSubmit = async () => {
		if (!isValidUserInput(isInvalid)) return;

		if (newPassword) {
			if (newPassword != passwordConfirmation) {
				isInvalid.passwordConfirmation = true;
				return;
			} else userData.new_password = newPassword;
		}

		try {
			let request = await api.put('/admin/minhaconta', userData);

			if (request.status == 200) document.location.href = '/admin';
		} catch (error) {
			errorFromServer = error.response.data;
		}
	};
</script>

<Card>
	<h1 class="text-center">Minha Conta</h1>
	<div class="mx-2">
		<p>Aqui você pode atualizar seus dados cadastrais ou apagar sua conta.</p>
		<p>Para qualquer atualização, é necessário fornecer sua senha no campo "Senha Atual".</p>
		<p>Caso deseje atualizar a senha, preencher "Nova senha" e "Confirmação de nova senha"</p>
	</div>
	<BlankForm>
		<div class="d-flex">
			<div class="me-1 w-100">
				<FormField
					name="first-name"
					label="Nome"
					placeholder="João"
					bind:value={userData.name}
					type="text"
					errorMsg="Seu nome deve possuir três ou mais letras"
					validation={/[\S ]{3,50}/}
					bind:isInvalid={isInvalid.name}
					required
				/>
			</div>
			<div class="ms-1 w-100">
				<FormField
					name="last-name"
					label="Sobrenome"
					placeholder="Silva"
					bind:value={userData.surname}
					type="text"
					errorMsg="Seu sobrenome deve possuir três ou mais letras"
					validation={/[\S ]{3,100}/}
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
			label="Senha"
			bind:value={userData.password}
			type="password"
			errorMsg="Sua senha deve possuir mais de oito dígitos"
			validation={/.{8,}/}
			bind:isInvalid={isInvalid.password}
			required
		/>
		<FormField
			name="new-passoword"
			label="Nova senha"
			bind:value={newPassword}
			errorMsg="Sua nova senha deve possuir mais de oito dígitos"
			validation={/.{8,}/}
			bind:isInvalid={isInvalid.newPassword}
			type="password"
		/>
		<FormField
			name="password-confirmation"
			label="Confirmação da senha"
			bind:value={passwordConfirmation}
			errorMsg="As senhas não são iguais"
			validation={/.*?/}
			type="password"
			bind:isInvalid={isInvalid.passwordConfirmation}
		/>

		{#if errorFromServer}
			<div class="mt-3 text-danger">
				<p>Ocorreu algum erro. Provavelmente isso ajude:</p>
				<p>{errorFromServer}</p>
			</div>
		{/if}

		<div class="mt-3 d-flex justify-content-end">
			<button type="submit" class="btn btn-primary" on:click={handleSubmit}>Atualizar</button>
		</div>
	</BlankForm>
</Card>

<style>
</style>

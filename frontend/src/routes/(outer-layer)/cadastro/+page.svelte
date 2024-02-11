<script>
	import FormField from '$lib/components/form-field.svelte';
	import BlankForm from '$lib/components/blank-form.svelte';
	import Card from '../../../lib/components/card.svelte';
	import { api } from '$lib/utils/api.js';

	let userData = {
		name: '',
		surname: '',
		email: '',
		password: '',
		birthday: '',
		phone_number: '',
		address: ''
	};

	let isInvalid = {
		name: false,
		surname: false,
		email: false,
		password: false,
		birthday: false,
		phone_number: false,
		address: false,
		passwordConfirmation: false
	};

	let passwordConfirmation;
	let errorFromServer;

	const isValidUserInput = () => {
		for (let key in isInvalid) {
			if (isInvalid[key]) return false;
		}
		return true;
	};

	const handleSubmit = async () => {
		if (!isValidUserInput()) return;

		if (userData.password != passwordConfirmation) {
			isInvalid.passwordConfirmation = true;
			return;
		}

		// go date format
		userData.birthday = `${userData.birthday}T00:00:00Z`;

		let request;
		try {
			request = await api.post('/cadastro', userData);
			// workaround for html "required"
			userData.birthday = userData.birthday.split('T')[0];

			if (request.status == 200) document.location.href = '/login';
		} catch (error) {
			errorFromServer = true;
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
					validation={/.{3,50}/}
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
					validation={/.{3,100}/}
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
			name="password-confirmation"
			label="Confirmação da senha"
			bind:value={passwordConfirmation}
			errorMsg="As senhas não são iguais"
			validation={/.*?/}
			type="password"
			bind:isInvalid={isInvalid.passwordConfirmation}
			required
		/>
		<FormField
			name="birthday"
			label="Data de nascimento &emsp;(mês/dia/ano)"
			bind:value={userData.birthday}
			type="date"
			errorMsg="Data de nascimento inválida"
			validation={/\d{1,4}-\d{1,2}-\d{1,2}/}
			bind:isInvalid={isInvalid.birthday}
			required
		/>
		<FormField
			name="phone-number"
			label="Número de telefone"
			bind:value={userData.phone_number}
			type="tel"
			placeholder="(41) 99999-9999"
			errorMsg="Número de telefone inválido. Utilizar formato (41) 99999-9999 ou (41) 3333-3333"
			validation={/(\(\d{2}\))?( )?\d{4,5}(-)?\d{4,5}/}
			bind:isInvalid={isInvalid.phone_number}
			required
		/>
		<FormField
			name="address"
			label="Endereço"
			bind:value={userData.address}
			type="text"
			placeholder="Rua dos Proletários, 1917"
			errorMsg="Seu endereço deve possuir mais de 5 dígitos"
			validation={/.{5,}/}
			bind:isInvalid={isInvalid.address}
			required
		/>

		{#if errorFromServer}
			<div class="mt-3 text-danger">
				<p>Ocorreu algum erro. Tente novamnente mais tarde.</p>
			</div>
		{/if}

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

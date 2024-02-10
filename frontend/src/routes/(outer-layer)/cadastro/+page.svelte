<script>
	import FormField from '$lib/components/form-field.svelte';
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

	let passwordConfirmation;

	const isValidUserInput = () => {
		let nameRegex = /.{3,50}/;
		let surnameRegex = /.{3,100}/;
		let emailRegex = /^[^\.\s][\w\-]+(\.[\w\-]+)*@([\w-]+\.)+[\w-]{2,}$/gm;
		let passwordRegex = /.{8,}/;
		let birthdayRegex = /\d{1,4}-\d{1,2}-\d{1,2}/;
		let phoneNumberRegex = /(\(\d{2}\))?( )?\d{4,5}(-)?\d{4,5}/;
		let addressRegex = /.{10,}/;

		if (!nameRegex.test(userData.name)) {
			console.log('name');
			return false;
		}

		if (!surnameRegex.test(userData.surname)) {
			console.log('surname');
			return false;
		}

		if (!emailRegex.test(userData.email)) {
			console.log('email');
			return false;
		}

		if (!passwordRegex.test(userData.password)) {
			console.log('passwrod');
			return false;
		}

		if (!birthdayRegex.test(userData.birthday)) {
			console.log('birthday');
			return false;
		}

		if (!phoneNumberRegex.test(userData.phone_number)) {
			console.log(userData.phone_number);
			console.log('phone');
			return false;
		}

		if (!addressRegex.test(userData.address)) {
			console.log('address');
			return false;
		}

		return true;
	};

	const handleSubmit = async () => {
		console.log(userData);
		if (!isValidUserInput()) return;

		let request = await api.post('/cadastro', userData);

		console.log(request);
	};
</script>

<Card>
	<h1>Cadastro</h1>
	<form>
		<div class="d-flex">
			<div class="me-1">
				<FormField
					name="first-name"
					label="Nome"
					placeholder="João"
					bind:value={userData.name}
					type="text"
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
			required
		/>
		<FormField
			name="password"
			label="Senha"
			bind:value={userData.password}
			type="password"
			required
		/>
		<FormField
			name="password-confirmation"
			label="Confirmação da senha"
			bind:value={passwordConfirmation}
			type="password"
			required
		/>
		<FormField
			name="birth-date"
			label="Data de nascimento"
			bind:value={userData.birthday}
			type="date"
			required
		/>
		<FormField
			name="phone-number"
			label="Número de telefone"
			bind:value={userData.phone_number}
			type="tel"
			placeholder="(41) 99999-9999"
			required
		/>
		<FormField
			name="address"
			label="Endereço"
			bind:value={userData.address}
			type="text"
			placeholder="Rua dos Proletários, 1917"
			required
		/>

		<div class="mt-3 d-flex justify-content-end">
			<button
				type="submit"
				class="btn btn-primary"
				on:click={() => {
					handleSubmit();
					//document.location.href = '/login';
				}}>Cadastrar</button
			>
		</div>
	</form>
</Card>

<style>
</style>

<script>
	import FormField from '$lib/components/form-field.svelte';
	import Card from '../../../lib/components/card.svelte';

	let userData = {
		firstName: '',
		lastName: '',
		email: '',
		password: '',
		birthDate: '',
		phoneNumber: '',
		address: ''
	};

	let passwordConfirmation;
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
					bind:value={userData.firstName}
					type="text"
					required
				/>
			</div>
			<div class="ms-1">
				<FormField
					name="last-name"
					label="Sobrenome"
					placeholder="Silva"
					bind:value={userData.lastName}
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
			bind:value={userData.birthDate}
			type="date"
			required
		/>
		<FormField
			name="phone-number"
			label="Número de telefone"
			bind:value={userData.phoneNumber}
			type="tel"
			placeholder="(41) 99999-9999"
			pattern="([0-9]{2}) [0-9]{5}-[0-9]{4}"
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
					document.location.href = '/login';
				}}>Cadastrar</button
			>
		</div>
	</form>
</Card>

<style>
</style>

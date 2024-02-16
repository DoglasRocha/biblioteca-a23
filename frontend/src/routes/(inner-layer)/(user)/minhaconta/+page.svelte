<script>
	import FormField from '$lib/components/form-field.svelte';
	import Card from '$lib/components/card.svelte';
	export let data;

	let userData = {
		name: '',
		surname: '',
		email: '',
		password: '',
		birthday: '',
		phone_number: '',
		address: ''
	};

	userData = data.userData;

	let newPassword, passwordConfirmation;
</script>

<Card>
	<h1>Minha Conta</h1>
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
			name="old-password"
			label="Senha Antiga"
			bind:value={userData.password}
			type="password"
			required
		/>
		<FormField
			name="new-passoword"
			label="Nova senha"
			bind:value={newPassword}
			type="password"
			required
		/>
		<FormField
			name="new-password-confirmation"
			label="Confirmação de nova senha"
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
				}}>Atualizar</button
			>
		</div>
	</form>
</Card>

<style>
</style>

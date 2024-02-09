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
			label="Senha Antiga"
			bind:value={userData.password}
			type="password"
			required
		/>
		<FormField
			name="password-confirmation"
			label="Nova senha"
			bind:value={passwordConfirmation}
			type="password"
			required
		/>
		<FormField
			name="password-confirmation"
			label="Confirmação de nova senha"
			bind:value={passwordConfirmation}
			type="password"
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

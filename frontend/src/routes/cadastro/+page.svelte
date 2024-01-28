<script>
	import FormField from '$lib/components/form-field.svelte';
	import Card from '../../lib/components/card.svelte';

	let userData = {
		firstName: '',
		lastName: '',
		email: '',
		password: '',
		birthDate: '',
		phoneNumber: '',
		address: ''
	};
</script>

<Card>
	<h1>Cadastro</h1>
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
</Card>

<style>
</style>

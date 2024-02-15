<script>
	import FormField from '$lib/components/form-field.svelte';
	import Card from '$lib/components/card.svelte';
	import BlankForm from '$lib/components/blank-form.svelte';

	let userData = {
		name: '',
		surname: '',
		email: '',
		password: ''
	};

	let passwordConfirmation;
</script>

<Card>
	<h1>Cadastro</h1>
	<BlankForm>
		<div class="d-flex">
			<div class="me-1">
				<FormField
					name="Name"
					label="Nome"
					placeholder="João"
					bind:value={userData.name}
					type="text"
					required
				/>
			</div>
			<div class="d-flex">
				<div class="me-1">
					<FormField
						name="surname"
						label="Surname"
						placeholder="Silva"
						bind:value={userData.surname}
						type="text"
						required
					/>
				</div>
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
	</BlankForm>
</Card>

<style>
</style>

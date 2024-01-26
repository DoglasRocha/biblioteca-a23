<script lang="ts">
	import FormField from '$lib/components/form-field.svelte';
</script>

<div class="card">
	<div class="card-body">
		<h1>Biblioteca do Cursinho A23</h1>

		<div class="mt-5">
			<h3>Login</h3>
		</div>

		<div class="mt-3">
			<label for="email" class="form-label">Endereço de Email</label>
			<input type="email" name="email" class="form-control" />
		</div>
		<div class="mt-3">
			<label for="senha" class="form-label">Senha</label>
			<input type="password" name="senha" class="form-control" />
		</div>
		<div class="buttons mt-3">
			<a class="btn btn-secondary" href="/cadastro"> Cadastrar </a>
			<button
				class="btn btn-primary"
				on:click={() => {
					window.location.href = '/';
				}}>Login</button
			>
		</div>
		<FormField></FormField>
	</div>
</div>

<style>
	.buttons {
		display: flex;
		justify-content: space-between;
	}
</style>

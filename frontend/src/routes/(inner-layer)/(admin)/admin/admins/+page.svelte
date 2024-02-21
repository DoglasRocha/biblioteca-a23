<script>
	import Card from '$lib/components/card.svelte';
	import Bold from '$lib/components/bold.svelte';
	export let data;

	let admins = data.admins,
		error = data.error;
</script>

<Card class="w-75">
	<h1 class="text-center">Administradores</h1>

	<div class="d-flex justify-content-center mt-5 px-5">
		<table class="table table-striped">
			<thead>
				<th>Nome</th>
			</thead>
			{#if error}
				<p>Ocorreu um erro. Talvez isso ajude:</p>
				<p>{error}</p>
			{:else}
				<tbody>
					{#each admins as admin}
						<tr>
							<td>
								<details>
									<summary>{admin.User.name} {admin.User.surname}</summary>

									<li><Bold>Email:</Bold> exemplo@email.com</li>
									<li><Bold>Status:</Bold> {admin.is_cleared ? 'Autorizado' : 'Não Autorizado'}</li>
									<div class="d-flex justify-content-end">
										{#if admin.is_cleared}
											<button type="button" class="btn btn-danger ms-1">Revogar </button>
										{:else}
											<button type="button" class="btn btn-success ms-1">Autorizar </button>
										{/if}
									</div>
								</details></td
							>
						</tr>
					{/each}
				</tbody>
			{/if}
		</table>
	</div>
</Card>

<style>
</style>

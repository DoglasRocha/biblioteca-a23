<script>
	import Card from '$lib/components/card.svelte';
	import Bold from '$lib/components/bold.svelte';
	import { api } from '$lib/utils/api.js';
	export let data;

	let admins = data.admins,
		error = data.error;

	const handleSubmit = async (action, adminId) => {
		try {
			let request = await api.patch(`/admin/${action}/${adminId}`);

			if (request.status == 200) document.location.href = document.location.href + '/';
		} catch (err) {
			error = err.response.data;
		}
	};
</script>

<Card class="w-75">
	<h1 class="text-center">Administradores</h1>

	<div class="d-flex justify-content-center mt-5 px-5">
		<table class="table table-striped">
			<thead>
				<th>Nome</th>
			</thead>
			<tbody>
				{#if admins}
					{#each admins as admin}
						<tr>
							<td>
								<details>
									<summary>{admin.User.name} {admin.User.surname}</summary>

									<li><Bold>Email:</Bold> {admin.User.email}</li>
									<li><Bold>Status:</Bold> {admin.is_cleared ? 'Autorizado' : 'Não Autorizado'}</li>
									<div class="d-flex justify-content-end">
										{#if admin.is_cleared}
											<button
												type="button"
												class="btn btn-danger ms-1"
												on:click={() => handleSubmit('revogar', admin.user_id)}
											>
												Revogar
											</button>
										{:else}
											<button
												type="button"
												class="btn btn-success ms-1"
												on:click={() => handleSubmit('autorizar', admin.user_id)}
											>
												Autorizar
											</button>
										{/if}
									</div>
								</details></td
							>
						</tr>
					{/each}
				{/if}
				{#if error}
					<tr>
						<td>
							<div class="text-danger">
								<p>Ocorreu um erro. Talvez isso ajude:</p>
								<p>{error}</p>
							</div>
						</td>
					</tr>
				{/if}
			</tbody>
		</table>
	</div>
</Card>

<style>
</style>

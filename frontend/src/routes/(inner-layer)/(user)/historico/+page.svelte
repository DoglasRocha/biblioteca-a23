<script>
	import Card from '$lib/components/card.svelte';
	import UserRequest from '$lib/components/user-request.svelte';

	export let data;

	let requests = data.requests,
		error = data.error;
</script>

<Card class="w-75">
	<h1 class="text-center">Solicitações</h1>
	<div class="d-flex justify-content-center mt-5 px-5">
		<table class="table table-striped">
			<thead>
				<th>Livro</th>
			</thead>
			<tbody>
				{#if requests.length}
					{#each requests as request}
						<tr>
							<td>
								<UserRequest {request} />
							</td>
						</tr>
					{/each}
				{:else}
					<tr>
						<td>
							<p>Não há solicitações em aberto!</p>
						</td>
					</tr>
				{/if}
			</tbody>
		</table>
	</div>
</Card>

<Card class="w-75">
	<h1 class="text-center">Empréstimos</h1>
	<div class="d-flex justify-content-center mt-5 px-5">
		<table class="table table-striped">
			<thead>
				<th>Livro</th>
			</thead>
			<tbody>
				{#if requests.length}
					{#each requests as request}
						<tr>
							<td>
								<UserRequest {request} />
							</td>
						</tr>
					{/each}
				{/if}
			</tbody>
		</table>
	</div>
</Card>

<style>
</style>

<script>
	import Card from '$lib/components/card.svelte';
	import Bold from '$lib/components/bold.svelte';
	import Loan from '$lib/components/loan.svelte';
	export let data;

	let loans = data.loans,
		error = data.error;
</script>

<Card class="w-75">
	<h1 class="text-center">Empréstimos ativos</h1>
	<div class="d-flex justify-content-center mt-5 px-5">
		<table class="table table-striped">
			<thead>
				<th>Empréstimos</th>
			</thead>
			<tbody>
				{#if loans.length}
					{#each loans as loan}
						<tr>
							<td>
								<Loan {loan}>
									<div class="d-flex justify-content-end">
										<button class="btn btn-primary">Confirmar devolução</button>
									</div>
								</Loan>
							</td>
						</tr>
					{/each}
				{:else}
					<tr><td>Não há empréstimos ativos.</td></tr>
				{/if}
				{#if error}
					<tr>
						<td class="text-danger">{error}</td>
					</tr>
				{/if}
			</tbody>
		</table>
	</div>
</Card>

<style>
</style>

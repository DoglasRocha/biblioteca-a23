<script>
	import Card from '$lib/components/card.svelte';
	import Bold from '$lib/components/bold.svelte';
	export let data;

	let loans = data.loans,
		error = data.error;

	console.log(loans);
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
								<details>
									<summary
										>{loan.Request.Book.name} - {loan.Request.Reader.User.name}
										{loan.Request.Reader.User.surname}</summary
									>
									<ul>
										<li>
											<Bold>Data de empréstimo:</Bold>
											{new Date(loan.start_date).toLocaleDateString('pt-br')}
										</li>
										<li>
											<Bold>Data de devolução:</Bold>
											{new Date(loan.return_date).toLocaleDateString('pt-br')}
										</li>
										<li>
											<Bold>Número da cópia do livro:</Bold>
											{loan.Copy.id}
										</li>
										<li>
											<Bold>Renovou:</Bold>
											{loan.has_renewed ? 'Sim' : 'Não'}
										</li>
										<li>
											<Bold>Está atrasado:</Bold>
											{new Date(loan.return_date) < Date.now() ? 'Sim' : 'Não'}
										</li>
									</ul>
									<div class="d-flex justify-content-end">
										<button class="btn btn-primary">Confirmar devolução</button>
									</div>
								</details>
							</td>
						</tr>
					{/each}
				{:else}
					<tr><td>Não há empréstimos ativos.</td></tr>
				{/if}
			</tbody>
		</table>
	</div>
</Card>

<style>
</style>

<script>
	import BaseNavbar from './base-navbar.svelte';
</script>

<BaseNavbar>
	<li class="nav-item">
		<a class="nav-link fw-bold" href="/admin">Home</a>
	</li>
	<li class="nav-item">
		<a class="nav-link fw-bold" href="/admin/buscar">Buscar livro</a>
	</li>
	<li class="nav-item">
		<a class="nav-link fw-bold" href="/admin/cadastrar">Cadastrar livro</a>
	</li>
	<li class="nav-item dropdown">
		<button class="nav-link dropdown-toggle fw-bold" data-bs-toggle="dropdown">Empréstimos</button>
		<ul class="dropdown-menu">
			<li class="px-1">
				<a class="nav-link fw-bold dropdown-item" href="/admin/emprestimos/solicitacoes"
					>Solicitações</a
				>
			</li>
			<li>
				<hr class="dropdown-divider" />
			</li>
			<li class="px-1">
				<a class="nav-link fw-bold dropdown-item" href="/admin/emprestimos/ativos">Ativos</a>
			</li>
			<li>
				<hr class="dropdown-divider" />
			</li>
			<li class="px-1">
				<a class="nav-link fw-bold dropdown-item" href="/admin/emprestimos/historico">Histórico</a>
			</li>
		</ul>
	</li>
</BaseNavbar>

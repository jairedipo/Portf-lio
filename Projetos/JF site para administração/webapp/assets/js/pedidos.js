$(document).on('load', carregarNomes("id_produto-1"));
$('#inserir-pedido').on('submit', inserirPedido);
$('#adicionar-produto').on('click', adicionarProduto);
$('#remover-produto').on('click', removerProduto);
$(document).on('change','.valor', calcularTaxa);
$(document).on('change','.valor-pedido', calcularValorPedido);
$(document).on('change','.valor-taxa', calcularValorTaxa);

// carregarNomes carrega os nomes dos produtos no campo para o autocomplete
function carregarNomes(elemento) {
  $.ajax({
    url: "/buscar-produtos",
    method: "GET",
    data: {}
  }).done(function(data) {
    autocomplete(document.getElementById(elemento), data);
  });
}

// inserirPedido insere um pedido na base
function inserirPedido(evento) {
    evento.preventDefault();
    var quantidadeProdutos = parseInt($('#quantidade-de-produtos').val());
    var i = 1;
    var inseridoComSucesso = true;

    while (i <= quantidadeProdutos) {

        $.ajax({
            url: "/pedidos",
            method: "POST",
            data: {
                id_pedido: $('#id_pedido').val(),
                dataDaCompra: $('#dataDaCompra').val(),
                id_produto: $('#id_produto-' + i).val(),
                quantidade: $('#quantidade-' + i).val(),
                valor: $('#valor-' + i).val(),
                valor_taxa: $('#valor_taxa-' + i).val(),
                estoque: $('#estoque').val(),
                embalagem: $('#embalagem').val(),
                plataforma: $('#plataforma').val(),
            }
        }).fail(function() {
            Swal.fire("Ops...", "Erro ao inserir pedido.", "error")
            .then(function() {
                const pedidoId = $('#id_pedido').val();
                $.ajax({
                    url: `/pedidos/${pedidoId}`,
                    method: "DELETE",
                })
            })
        })

        i = i + 1;
    }

    if (inseridoComSucesso) {
        Swal.fire("Sucesso", "Pedido cadastrado com sucesso", "success").then( function() {
            window.location = "/pedidos/cadastrar-pedidos";
        })
    }
}

// calcularTaxa calcula e preenche automaticamente a taxa baseada no valor do produto informado
function calcularTaxa(evento) {
    evento.preventDefault();

    const campoTaxa = document.querySelector('#valor_taxa'+evento.target.id.substr(5));
    if (evento.target.value > 6) {
      campoTaxa.value = (evento.target.value * 0.2 + 3.0).toFixed(2);
    }
    else {
      campoTaxa.value = (evento.target.value * 0.7).toFixed(2);
    }
    calcularValorTaxa()
}

// calcularValorPedido calcula o valor total do pedido
function calcularValorPedido(evento) {
  var campoQuantidadeProdutos = document.querySelector('#quantidade-de-produtos');
  var count;
  var valorTotal = 0;
  for (count = 1; count <= campoQuantidadeProdutos.value; count++) {
    var quantidade = document.querySelector("#quantidade-" + count).value;
    var valor = document.querySelector("#valor-" + count).value;
    valorTotal = valorTotal + (quantidade * valor);
  }
  document.querySelector('#valor-total-pedido').textContent = "R$ " + valorTotal.toFixed(2);
  calcularValorRecebido();
}

// calcularValorTaxa calcula o valor total das taxas do pedido
function calcularValorTaxa(evento) {
  var campoQuantidadeProdutos = document.querySelector('#quantidade-de-produtos');
  var count;
  var valorTotalTaxa = 0;
  for (count = 1; count <= campoQuantidadeProdutos.value; count++) {
    var quantidade = document.querySelector("#quantidade-" + count).value;
    var valorTaxa = Number(document.querySelector("#valor_taxa-" + count).value);
    valorTotalTaxa = valorTotalTaxa + (valorTaxa * quantidade);
  }
  document.querySelector('#valor-total-taxa').textContent = "-R$ " + valorTotalTaxa.toFixed(2);
  calcularValorRecebido();
}

// calcularValorRecebido calcula o valor a receber do pedido
function calcularValorRecebido(evento) {
  var valor = parseFloat(document.querySelector('#valor-total-pedido').textContent.substr(3));
  var taxa = parseFloat(document.querySelector('#valor-total-taxa').textContent.substr(4));
  var valorRecebido = valor - taxa;
  document.querySelector('#valor-recebido').textContent = "R$ " + valorRecebido.toFixed(2);
}

// adicionarProduto adiciona campos para inclusão de mais um produto no pedido
function adicionarProduto(evento) {
  var templateProduto = $(".template-produto");
  var textoTemplateProduto = templateProduto.html().toString();

  var campoQuantidadeProdutos = document.querySelector('#quantidade-de-produtos');
  campoQuantidadeProdutos.value = parseInt(campoQuantidadeProdutos.value) + 1;
  textoTemplateProduto = textoTemplateProduto.replace(/-1/g, '-' + campoQuantidadeProdutos.value);
  textoTemplateProduto = textoTemplateProduto.replace(/ 1/g, ' ' + campoQuantidadeProdutos.value);

  var div = document.createElement('div');
  div.classList.add('template-produto')
  div.innerHTML = textoTemplateProduto;
  document.body.querySelector('#lista-produtos').append(div);

  var elemento = "id_produto-" + campoQuantidadeProdutos.value;
  carregarNomes(elemento)
}

// removerProduto remove os campos do último produto
function removerProduto(evento) {
  var listaProdutos = document.querySelector('#lista-produtos');
  var ultimoProduto = listaProdutos.lastChild;
  var campoQuantidadeProdutos = document.querySelector('#quantidade-de-produtos');

  if (ultimoProduto.parentNode && campoQuantidadeProdutos.value > 1) {
    ultimoProduto.parentNode.removeChild(ultimoProduto);
    campoQuantidadeProdutos.value = parseInt(campoQuantidadeProdutos.value) - 1;
  }
}

// https://www.w3schools.com/howto/howto_js_autocomplete.asp
function autocomplete(inp, arr) {
  /*the autocomplete function takes two arguments,
  the text field element and an array of possible autocompleted values:*/
  var currentFocus;
  /*execute a function when someone writes in the text field:*/
  inp.addEventListener("input", function(e) {
      var a, b, i, val = this.value;
      /*close any already open lists of autocompleted values*/
      closeAllLists();
      if (!val) { return false;}
      currentFocus = -1;
      /*create a DIV element that will contain the items (values):*/
      a = document.createElement("DIV");
      a.setAttribute("id", this.id + "autocomplete-list");
      a.setAttribute("class", "autocomplete-items");
      /*append the DIV element as a child of the autocomplete container:*/
      this.parentNode.appendChild(a);
      /*for each item in the array...*/
      for (i = 0; i < arr.length; i++) {
        /*check if the item starts with the same letters as the text field value:*/
        if (arr[i].substr(0, val.length).toUpperCase() == val.toUpperCase()) {
          /*create a DIV element for each matching element:*/
          b = document.createElement("DIV");
          /*make the matching letters bold:*/
          b.innerHTML = "<strong>" + arr[i].substr(0, val.length) + "</strong>";
          b.innerHTML += arr[i].substr(val.length);
          /*insert a input field that will hold the current array item's value:*/
          b.innerHTML += "<input type='hidden' value='" + arr[i] + "'>";
          /*execute a function when someone clicks on the item value (DIV element):*/
          b.addEventListener("click", function(e) {
              /*insert the value for the autocomplete text field:*/
              inp.value = this.getElementsByTagName("input")[0].value;
              /*close the list of autocompleted values,
              (or any other open lists of autocompleted values:*/
              closeAllLists();
          });
          a.appendChild(b);
        }
      }
  });
  /*execute a function presses a key on the keyboard:*/
  inp.addEventListener("keydown", function(e) {
      var x = document.getElementById(this.id + "autocomplete-list");
      if (x) x = x.getElementsByTagName("div");
      if (e.keyCode == 40) {
        /*If the arrow DOWN key is pressed,
        increase the currentFocus variable:*/
        currentFocus++;
        /*and and make the current item more visible:*/
        addActive(x);
      } else if (e.keyCode == 38) { //up
        /*If the arrow UP key is pressed,
        decrease the currentFocus variable:*/
        currentFocus--;
        /*and and make the current item more visible:*/
        addActive(x);
      } else if (e.keyCode == 13) {
        /*If the ENTER key is pressed, prevent the form from being submitted,*/
        e.preventDefault();
        if (currentFocus > -1) {
          /*and simulate a click on the "active" item:*/
          if (x) x[currentFocus].click();
        }
      }
  });
  function addActive(x) {
    /*a function to classify an item as "active":*/
    if (!x) return false;
    /*start by removing the "active" class on all items:*/
    removeActive(x);
    if (currentFocus >= x.length) currentFocus = 0;
    if (currentFocus < 0) currentFocus = (x.length - 1);
    /*add class "autocomplete-active":*/
    x[currentFocus].classList.add("autocomplete-active");
  }
  function removeActive(x) {
    /*a function to remove the "active" class from all autocomplete items:*/
    for (var i = 0; i < x.length; i++) {
      x[i].classList.remove("autocomplete-active");
    }
  }
  function closeAllLists(elmnt) {
    /*close all autocomplete lists in the document,
    except the one passed as an argument:*/
    var x = document.getElementsByClassName("autocomplete-items");
    for (var i = 0; i < x.length; i++) {
      if (elmnt != x[i] && elmnt != inp) {
        x[i].parentNode.removeChild(x[i]);
      }
    }
  }
  /*execute a function when someone clicks in the document:*/
  document.addEventListener("click", function (e) {
      closeAllLists(e.target);
  });
}
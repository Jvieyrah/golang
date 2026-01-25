package com.eliasnogueira;

import org.assertj.core.api.Assertions;
import org.junit.jupiter.api.Test;

import java.math.BigDecimal;

public class DescontosServiceTest {

    DescontoService descontosService = new DescontoService();

    @Test
    void deveAplicarDescontoQuandoMaiorQue100() {
        BigDecimal valor = new BigDecimal("150.00");
        BigDecimal valorComDesconto = descontosService.aplicarDesconto(valor);
        Assertions.assertThat(valorComDesconto).isEqualByComparingTo("135.00");
    }

    @Test
    void naoDeveAplicarDescontoQuandoMenorQue100() {
        BigDecimal valor = new BigDecimal("99.99");
        BigDecimal valorComDesconto = descontosService.aplicarDesconto(valor);
        Assertions.assertThat(valorComDesconto).isEqualByComparingTo(valor);

    }


}


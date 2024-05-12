+++
title = 'tables'
date = 2024-05-12
+++

Tables are a great way to display data in a structured format. They can be used to display a variety of data types, from simple text to more complex data sets. In this post, we'll take a look at how to create tables using HTML and CSS.

<div class="table-scroller">
  <table class="table max-w-[800px]">
    <caption class="table-caption">A list of your recent invoices</caption>
    <thead>
      <tr class="table-tr">
        <th class="table-th">Invoice</th>
        <th class="table-th">Status</th>
        <th class="table-th">Method</th>
        <th class="table-th text-right">Amount</th>
      </tr>
    </thead>
    <tbody>
    <tr class="table-tr">
      <td class="table-td">INV001</td>
      <td class="table-td">Paid</td>
      <td class="table-td">Credit Card</td>
      <td class="table-td text-right">£250.00</td>
    </tr>
    <tr class="table-tr">
      <td class="table-td">INV002</td>
      <td class="table-td">Unpaid</td>
      <td class="table-td">Paypal</td>
      <td class="table-td text-right">£100.00</td>
    </tr>
    <tr class="table-tr">
      <td class="table-td">INV003</td>
      <td class="table-td">Paid</td>
      <td class="table-td">Bank Transfer</td>
      <td class="table-td text-right">£500.00</td>
    </tr>
    <tr class="table-tr">
      <td class="table-td">INV004</td>
      <td class="table-td">Unpaid</td>
      <td class="table-td">Credit Card</td>
      <td class="table-td text-right">£150.00</td>
    </tr>
    <tr class="table-tr">
      <td class="table-td">INV005</td>
      <td class="table-td">Paid</td>
      <td class="table-td">Bank Transfer</td>
      <td class="table-td text-right">£300.00</td>
    </tr>
    <tr class="table-tr">
      <td class="table-td">INV006</td>
      <td class="table-td">Unpaid</td>
      <td class="table-td">Paypal</td>
      <td class="table-td text-right">£200.00</td>
    </tr>
    <tr class="table-tr">
      <td class="table-td">INV007</td>
      <td class="table-td">Paid</td>
      <td class="table-td">Credit Card</td>
      <td class="table-td text-right">£350.00</td>
    </tr>
    </tbody>
    <tfoot>
      <tr class="table-tr table-tfoot-tr">
        <td class="table-td" colspan="3">Total</td>
        <td class="table-td text-right">£1850.00</td>
      </tr>
    </tfoot>
  </table>
</div>
+++
title = 'forms'
date = 2024-05-12
+++

Forms are a key part of any website. They allow users to input data and submit it to the server. This data can then be used to create new records, update existing records, or perform other actions.

<div class="space-y-6">
  <div class="space-y-2">
    <label class="input-label">Email</label>
    <input class="input" placeholder="email address" type="email" />
    <div class="input-help">This is your help text</div>
  </div>
  <div class="space-y-2">
    <label class="input-label">File</label>
    <input class="input" type="file" />
  </div>
  <div class="space-y-2">
    <label class="input-label">Name</label>
    <input class="input" placeholder="your name" type="text" />
    <div class="input-error">This field is required</div>
  </div>
  <div class="space-y-2">
    <label class="input-label">Country</label>
    <select class="select">
      <option selected>Choose a country</option>
      <option value="US">United States</option>
      <option value="CA">Canada</option>
      <option value="FR">France</option>
      <option value="DE">Germany</option>
    </select>
  </div>
  <div class="space-y-2">
    <label class="input-label">Message</label>
    <textarea class="input min-h-[80px]" placeholder="add your message"></textarea>
  </div>
  <div class="flex w-full max-w-sm items-center space-x-2">
    <input class="input" placeholder="email address" type="email" />
    <button class="button button-primary">button</button>
  </div>
</div>

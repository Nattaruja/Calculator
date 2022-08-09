<template>

<div class="vertical-nav bg-purple" id="sidebar">
    <div class="py-4 px-3 mb-4">
      <div class="media d-flex align-items-center"><img src="../img/gpslogo.png" alt="..." width="65" class="mr-3 rounded-circle img-thumbnail shadow-sm">
        <div class="media-body">
          <h5 class="m-0">Good Packaging Supply</h5>
        </div>
      </div>
    </div>
  
    <p class="font-weight-bold text-uppercase px-3 small pb-4 mb-0">Main</p>
  
    <ul class="nav flex-column mb-0">
      <li class="nav-item">
        <a href="/v1/item/items" class="nav-link bg-purple2">
          Item
        </a>
      </li>
      <li class="nav-item">
        <a href="/v1/customer/customers" class="nav-link bg-purple2">
          Customer
        </a>
      </li>
    </ul>
</div>

</template>
<script>
import '../App.css';
export default {
 name: 'MainSidebar'
}
</script>
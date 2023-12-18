import { useEffect, useState } from 'react';
import 'bootstrap/dist/css/bootstrap.css'
import ItemForm from './itemForm';
export default function Home() {
  const [addItem, setAddItem] = useState(true)
  const [itemForm, setItemForm] = useState([])
  const [listProduct, setListProduct] = useState([])
  const [customerName, setCustomerName] = useState('')  
  
  const addItems = ()=>{
    setAddItem(true)
  } 

  const cancel = () => {
    setCustomerName('')
    setItemForm([])
    setAddItem(true)
  }  
  
  useEffect(() => {
    const fetchData = async() => {
      const resp = await fetch('/api/product')
      const response = await resp.json()
      setListProduct(response.data)
    }
    fetchData()
  },[])
  
  const handleChangeProduct = (event) => {
    event.preventDefault()
    let value = event.target.value.split('-')
    let newItemForm = itemForm.map(dt => {
      if (dt.id == value[1]){
        let selectedData = listProduct.filter(product => {
          return product.id == value[0]
        })[0]
        dt.product_id = selectedData.id
        dt.price = selectedData.price
        dt.total = dt.price * dt.qty
      }
      return dt
    }) 
    setItemForm(newItemForm)
  }

  const handleChangeQty = (event) => {
    event.preventDefault()
    let value = parseInt(event.target.value)

    let newItemForm = itemForm.map(dt => {
      if (dt.id == event.target.getAttribute('data-key')){
        dt.qty = value
        dt.total = dt.price * dt.qty
      }
      return dt
    }) 
    setItemForm(newItemForm)
  }

  const handleChangeName = (event) => {
    event.preventDefault()
    setCustomerName(event.target.value)    
  }

  const onSubmit = () => {

    const dataProduct = itemForm.map((dt) => {
      return {
        product_id: dt.product_id,
        qty: dt.qty
      }
    })
    const data = {
      customer_name: customerName,
      product_list: dataProduct
    }
    fetch('/api/transaction', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
    })
        .then(async(response) => {
          const dataRes = await response.json()
          if (dataRes.data)
            alert('Your transaction code is '+ dataRes.data)
          else if (dataRes.message)
            alert(dataRes.message)
        })
        .then((data) => console.log(data))
  }

  useEffect(()=>{
    if (addItem && listProduct.length>0) {
      const selectedData = listProduct[0]
      setItemForm([...itemForm, {id: itemForm.length + 1, product_id: selectedData.id, price: selectedData.price, qty: 1, total: selectedData.price}])
      setAddItem(false)
    }
  },[addItem, listProduct])
  
  return (
    <div className='col-md-10 py-4 px-5'>
      <form>
        <div className='my-3 py-2 px-2 border border-dark'>
          <label>Nama Customer</label>
          <input className='form-control' type="text" name="customer_name" value={customerName} onChange={handleChangeName} />
        </div>
        <div className='col-md-12 my-3 px-2 border border-dark'>
          {itemForm.length > 0 ? itemForm.map(item => <ItemForm product={listProduct} data={item} setQty={handleChangeQty} setItemValue={handleChangeProduct}/>) : <></>}
          <div className='text-end pt-3 pb-2'>
            <button className='btn btn-primary' type='button' onClick={()=> addItems()}>Add Items</button>
          </div>
        </div>
        <div className='col-md-12 text-end pt-2'>
          <button className='btn btn-warning me-2' type='button' onClick={()=> cancel()}>Cancel</button>
          <button className='btn btn-success' type='button' onClick={()=> onSubmit()}>Submit</button>
        </div>
      </form>
    </div>
  );
}

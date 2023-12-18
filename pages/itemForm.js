export default function ItemForm(prop){
    return(
        <div className='row pt-2' key={'item-'+prop.data?.id}>
            <div className="col-md-3">
                <label>Nama Barang</label>
                <select className='form-control' onChange={prop.setItemValue}>
                    {
                        prop.product?.length > 0 ?
                        prop.product.map(dt => {
                            return <option key={dt?.id+'-'+prop.data?.id} value={dt?.id+'-'+prop.data?.id} data={dt}>{dt.name}</option>
                        }) : <></>
                    }
                </select>
            </div>
            <div className="col-md-3">
                <label>Harga Satuan</label>
                <input className='form-control' value={prop.data?.price || 0} disabled/>
            </div>
            <div className="col-md-3">
                <label >Qty</label>
                <input className='form-control' value={prop.data?.qty || 0} data-key={prop.data?.id} type="number" onChange={prop.setQty}/>
            </div>
            <div className="col-md-3">
                <label>Total</label>
                <input className='form-control' value={prop.data?.total || 0} disabled/>
            </div>
        </div>
        
    )

}
// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import JscCheckers from './jsc.checkers'
import JscJsc from './jsc.jsc'


export default { 
  JscCheckers: load(JscCheckers, 'jsc.checkers'),
  JscJsc: load(JscJsc, 'jsc.jsc'),
  
}


function load(mod, fullns) {
    return function init(store) {        
        if (store.hasModule([fullns])) {
            throw new Error('Duplicate module name detected: '+ fullns)
        }else{
            store.registerModule([fullns], mod)
            store.subscribe((mutation) => {
                if (mutation.type == 'common/env/INITIALIZE_WS_COMPLETE') {
                    store.dispatch(fullns+ '/init', null, {
                        root: true
                    })
                }
            })
        }
    }
}
package components

import react.*
import react.dom.*

@JsModule("react-select/dist/react-select")
external val reactSelect: RClass<SelectProps>

external interface Option {
    var value: String
    var label: String
}

external interface SelectProps : RProps {
    var name: String
    var options: Array<Option>
    var value: String?
    var onInputChange: (String) -> Unit
    var onChange: (Option) -> Unit
    var isLoading: Boolean
    var backspaceRemoves: Boolean
}

class CustomSelect(props: SelectProps) : RComponent<SelectProps, RState>(props) {

    override fun RBuilder.render() {        
        div {            
            reactSelect {
                attrs {
                    name=props.name
                    options=props.options
                    value=props.value
                    onInputChange=props.onInputChange
                    onChange=props.onChange
                    isLoading=props.isLoading
                    backspaceRemoves=props.backspaceRemoves
                }
            }                    
        }
    }
}

fun RBuilder.customSelect(
    name: String,
    options: Array<Option>,
    value: String?,
    onInputChange: (String) -> Unit, 
    onChange: (Option) -> Unit,
    isLoading: Boolean = false,
    backspaceRemoves: Boolean = true    
) = child(CustomSelect::class) {
                attrs.name=name
                attrs.options=options
                attrs.value=value
                attrs.onInputChange=onInputChange
                attrs.onChange=onChange
                attrs.isLoading=isLoading
                attrs.backspaceRemoves=backspaceRemoves    
}

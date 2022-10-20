import { PaymentMethodsInterface } from "./IPaymentmethod";
import { DeliveryTypesInterface } from "./IDeliverytype";
import { OrdersInterface } from "./IOrder";

export interface PaymentInterface {
  ID: number,
  Phone: string,
  Price: Float32Array
  PaymentTime: Date,
  DeliveryTypeID: number,
  DeliveryType: DeliveryTypesInterface,
  PaymentMethodID: number,
  PaymentMethod: PaymentMethodsInterface,
  OrderID: number,
  Order: OrdersInterface,
}
import { createStyles, makeStyles, Theme } from "@material-ui/core/styles";
import Container from "@material-ui/core/Container";

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    container: {
      marginTop: theme.spacing(2),
    },
    table: {
      minWidth: 650,
    },
    tableSpace: {
      marginTop: 20,
    },
  })
);

function Home() {
  const classes = useStyles();

  return (
    <div>
      <Container className={classes.container} maxWidth="md">
        <h1 style={{ textAlign: "center" }}>ระบบ Farm  Mart</h1>
        <h4>Requirements</h4>
        <p>
          ระบบ Farm mart เป็นระบบที่ผู้ใช้ระบบซึ่งเป็นสมาชิกสามารถ login เข้ามาเพื่อใช้งานระบบ Farm mart
          เพื่อทำการค้นหาสินค้า และเลือกซื้อสินค้าที่สนใจได้ เมื่อสมาชิกกดสั่งซื้อสินค้าในระบบเรียบร้อยแล้ว
          ระบบจะทำการออกใบสั่งซื้อสินค้าที่มีรายการสินค้าทั้งหมด และกำหนดเลขที่ใบสั่งซื้อพร้อมทั้งรวมยอดเงินทั้งหมดของรายการสินค้านั้นไว้
          เพื่อให้สมาชิกทำการชำระเงิน สมาชิกสามารถเลือกช่องทางการชำระเงิน และการจัดส่งที่ตนเองสะดวกได้
          เมื่อสมาชิกทำการชำระเงินเรียบร้อยแล้วระบบจะขึ้นข้อความว่าการชำระเงินเสร็จสมบูรณ์แล้วระบบจะทำการบันทึกข้อมูลการชำระเงิน
          และเตรียมจัดส่งสินค้าให้กับสมาชิก
        </p>
      </Container>
    </div>
  );
}
export default Home;